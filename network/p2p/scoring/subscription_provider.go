package scoring

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/rs/zerolog"
	"go.uber.org/atomic"

	"github.com/onflow/flow-go/module"
	"github.com/onflow/flow-go/module/component"
	"github.com/onflow/flow-go/module/irrecoverable"
	"github.com/onflow/flow-go/network/p2p"
	"github.com/onflow/flow-go/network/p2p/p2pconf"
	"github.com/onflow/flow-go/network/p2p/p2plogging"
)

// SubscriptionProvider provides a list of topics a peer is subscribed to.
type SubscriptionProvider struct {
	component.Component
	logger              zerolog.Logger
	topicProviderOracle func() p2p.TopicProvider

	// allTopics is a list of all topics in the pubsub network
	// TODO: we should add an expiry time to this cache and clean up the cache periodically
	// to avoid leakage of stale topics.
	cache SubscriptionCache

	// idProvider translates the peer ids to flow ids.
	idProvider module.IdentityProvider

	// allTopics is a list of all topics in the pubsub network that this node is subscribed to.
	allTopicsUpdate         atomic.Bool   // whether a goroutine is already updating the list of topics
	allTopicsUpdateInterval time.Duration // the interval for updating the list of topics in the pubsub network that this node has subscribed to.
}

type SubscriptionProviderConfig struct {
	Logger              zerolog.Logger           `validate:"required"`
	TopicProviderOracle func() p2p.TopicProvider `validate:"required"`
	IdProvider          module.IdentityProvider  `validate:"required"`
	Params              *p2pconf.SubscriptionProviderParameters
}

var _ p2p.SubscriptionProvider = (*SubscriptionProvider)(nil)

func NewSubscriptionProvider(cfg *SubscriptionProviderConfig) (*SubscriptionProvider, error) {
	if err := validator.New().Struct(cfg); err != nil {
		return nil, fmt.Errorf("invalid subscription provider config: %w", err)
	}

	p := &SubscriptionProvider{
		logger:                  cfg.Logger.With().Str("module", "subscription_provider").Logger(),
		topicProviderOracle:     cfg.TopicProviderOracle,
		allTopicsUpdateInterval: cfg.Params.SubscriptionUpdateInterval,
	}

	builder := component.NewComponentManagerBuilder()
	p.Component = builder.AddWorker(
		func(ctx irrecoverable.SignalerContext, ready component.ReadyFunc) {
			ready()
			p.logger.Debug().Msg("subscription provider started; starting update topics loop")
			p.updateTopicsLoop(ctx)

			<-ctx.Done()
			p.logger.Debug().Msg("subscription provider stopped; stopping update topics loop")
		}).Build()

	return p, nil
}

func (s *SubscriptionProvider) updateTopicsLoop(ctx irrecoverable.SignalerContext) {
	ticker := time.NewTicker(s.allTopicsUpdateInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.updateTopics(ctx)
		}
	}
}

// updateTopics returns all the topics in the pubsub network that this node (peer) has subscribed to.
// Note that this method always returns the cached version of the subscribed topics while querying the
// pubsub network for the list of topics in a goroutine. Hence, the first call to this method always returns an empty
// list.
func (s *SubscriptionProvider) updateTopics(ctx irrecoverable.SignalerContext) {
	if updateInProgress := s.allTopicsUpdate.CompareAndSwap(false, true); updateInProgress {
		// another goroutine is already updating the list of topics
		return
	}

	// start of critical section; protected by updateInProgress atomic flag
	allTopics := s.topicProviderOracle().GetTopics()
	s.logger.Trace().Msgf("all topics updated: %v", allTopics)

	// increments the update cycle of the cache; so that the previous cache entries are invalidated upon a read or write.
	s.cache.MoveToNextUpdateCycle()
	for _, topic := range allTopics {
		peers := s.topicProviderOracle().ListPeers(topic)

		if _, authorized := s.idProvider.ByPeerID(peers[0]); !authorized {
			// peer is not authorized (staked); hence it does not have a valid role in the network; and
			// we skip the topic update for this peer (also avoiding sybil attacks on the cache).
			continue
		}

		for _, p := range peers {
			updatedTopics, err := s.cache.AddTopicForPeer(p, topic)
			if err != nil {
				// this is an irrecoverable error; hence, we crash the node.
				ctx.Throw(fmt.Errorf("failed to update topics for peer %s: %w", p, err))
			}
			s.logger.Debug().
				Str("remote_peer_id", p2plogging.PeerId(p)).
				Strs("updated_topics", updatedTopics).
				Msg("updated topics for peer")
		}
	}

	// remove the update flag; end of critical section
	s.allTopicsUpdate.Store(false)
}

// GetSubscribedTopics returns all the subscriptions of a peer within the pubsub network.
func (s *SubscriptionProvider) GetSubscribedTopics(pid peer.ID) []string {
	topics, ok := s.cache.GetSubscribedTopics(pid)
	if !ok {
		s.logger.Trace().Str("peer_id", p2plogging.PeerId(pid)).Msg("no topics found for peer")
		return nil
	}
	return topics
}

package pubsub

import (
	"github.com/onflow/flow-go/consensus/hotstuff"
)

// Distributor distributes notifications to a list of subscribers (event consumers).
//
// It allows thread-safe subscription of multiple consumers to events.
type Distributor struct {
	FollowerDistributor
	CommunicatorDistributor
	ParticipantDistributor
}

var _ hotstuff.Consumer = (*Distributor)(nil)

func NewDistributor() *Distributor {
	return &Distributor{}
}

// AddConsumer adds an event consumer to the Distributor
func (p *Distributor) AddConsumer(consumer hotstuff.Consumer) {
	p.FollowerDistributor.AddFollowerConsumer(consumer)
	p.CommunicatorDistributor.AddCommunicatorConsumer(consumer)
	p.ParticipantDistributor.AddParticipantConsumer(consumer)
}

// FollowerDistributor ingests consensus follower events and distributes it to subscribers.
type FollowerDistributor struct {
	ProtocolViolationDistributor
	FinalizationDistributor
}

var _ hotstuff.FollowerConsumer = (*FollowerDistributor)(nil)

func NewFollowerDistributor() *FollowerDistributor {
	return &FollowerDistributor{}
}

func (d *FollowerDistributor) AddFollowerConsumer(consumer hotstuff.FollowerConsumer) {
	d.FinalizationDistributor.AddFinalizationConsumer(consumer)
	d.ProtocolViolationDistributor.AddProtocolViolationConsumer(consumer)
}

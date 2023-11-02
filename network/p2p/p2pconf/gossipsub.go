package p2pconf

import (
	"time"
)

// ResourceManagerConfig returns the resource manager configuration for the libp2p node.
// The resource manager is used to limit the number of open connections and streams (as well as any other resources
// used by libp2p) for each peer.
type ResourceManagerConfig struct {
	Override                  ResourceManagerOverrideScope `mapstructure:"libp2p-resource-limit-override"`        // override limits for specific peers, protocols, etc.
	MemoryLimitRatio          float64                      `mapstructure:"libp2p-memory-limit-ratio"`             // maximum allowed fraction of memory to be allocated by the libp2p resources in (0,1]
	FileDescriptorsRatio      float64                      `mapstructure:"libp2p-file-descriptors-ratio"`         // maximum allowed fraction of file descriptors to be allocated by the libp2p resources in (0,1]
	PeerBaseLimitConnsInbound int                          `mapstructure:"libp2p-peer-base-limits-conns-inbound"` // the maximum amount of allowed inbound connections per peer
}

type ResourceManagerOverrideScope struct {
	// System is the limit for the resource at the entire system.
	// For a specific limit, the system-wide dictates the maximum allowed value across all peers and protocols at the entire node level.
	System ResourceManagerOverrideLimit `mapstructure:"system"`

	// Transient is the limit for the resource at the transient scope. Transient limits are used for resources that have not fully established and are under negotiation.
	Transient ResourceManagerOverrideLimit `mapstructure:"transient"`

	// Protocol is the limit for the resource at the protocol scope, e.g., DHT, GossipSub, etc. It dictates the maximum allowed resource across all peers for that protocol.
	Protocol ResourceManagerOverrideLimit `mapstructure:"protocol"`

	// Peer is the limit for the resource at the peer scope. It dictates the maximum allowed resource for a specific peer.
	Peer ResourceManagerOverrideLimit `mapstructure:"peer"`

	// Connection is the limit for the resource for a pair of (peer, protocol), e.g., (peer1, DHT), (peer1, GossipSub), etc. It dictates the maximum allowed resource for a protocol and a peer.
	PeerProtocol ResourceManagerOverrideLimit `mapstructure:"peer-protocol"`
}

// ResourceManagerOverrideLimit is the configuration for the resource manager override limit at a certain scope.
// Any value that is not set will be ignored and the default value will be used.
type ResourceManagerOverrideLimit struct {
	// System is the limit for the resource at the entire system. if not set, the default value will be used.
	// For a specific limit, the system-wide dictates the maximum allowed value across all peers and protocols at the entire node scope.
	StreamsInbound int `validate:"gte=0" mapstructure:"stream-inbound"`

	// StreamsOutbound is the max number of outbound streams allowed, at the resource scope.
	StreamsOutbound int `validate:"gte=0" mapstructure:"stream-outbound"`

	// ConnectionsInbound is the max number of inbound connections allowed, at the resource scope.
	ConnectionsInbound int `validate:"gte=0" mapstructure:"connection-inbound"`

	// ConnectionsOutbound is the max number of outbound connections allowed, at the resource scope.
	ConnectionsOutbound int `validate:"gte=0" mapstructure:"connection-outbound"`

	// FD is the max number of file descriptors allowed, at the resource scope.
	FD int `validate:"gte=0" mapstructure:"fd"`

	// Memory is the max amount of memory allowed (bytes), at the resource scope.
	Memory int `validate:"gte=0" mapstructure:"memory-bytes"`
}

// GossipSubConfig is the configuration for the GossipSub pubsub implementation.
type GossipSubConfig struct {
	// GossipSubRPCInspectorsConfig configuration for all gossipsub RPC control message inspectors.
	GossipSubRPCInspectorsConfig `mapstructure:",squash"`

	// GossipSubTracerConfig is the configuration for the gossipsub tracer. GossipSub tracer is used to trace the local mesh events and peer scores.
	GossipSubTracerConfig `mapstructure:",squash"`

	// PeerScoring is whether to enable GossipSub peer scoring.
	PeerScoring bool `mapstructure:"gossipsub-peer-scoring-enabled"`
}

// GossipSubTracerConfig is the config for the gossipsub tracer. GossipSub tracer is used to trace the local mesh events and peer scores.
type GossipSubTracerConfig struct {
	// LocalMeshLogInterval is the interval at which the local mesh is logged.
	LocalMeshLogInterval time.Duration `validate:"gt=0s" mapstructure:"gossipsub-local-mesh-logging-interval"`
	// ScoreTracerInterval is the interval at which the score tracer logs the peer scores.
	ScoreTracerInterval time.Duration `validate:"gt=0s" mapstructure:"gossipsub-score-tracer-interval"`
	// RPCSentTrackerCacheSize cache size of the rpc sent tracker used by the gossipsub mesh tracer.
	RPCSentTrackerCacheSize uint32 `validate:"gt=0" mapstructure:"gossipsub-rpc-sent-tracker-cache-size"`
	// RPCSentTrackerQueueCacheSize cache size of the rpc sent tracker queue used for async tracking.
	RPCSentTrackerQueueCacheSize uint32 `validate:"gt=0" mapstructure:"gossipsub-rpc-sent-tracker-queue-cache-size"`
	// RpcSentTrackerNumOfWorkers number of workers for rpc sent tracker worker pool.
	RpcSentTrackerNumOfWorkers int `validate:"gt=0" mapstructure:"gossipsub-rpc-sent-tracker-workers"`
}

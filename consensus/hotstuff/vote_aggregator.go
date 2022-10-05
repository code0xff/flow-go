package hotstuff

import (
	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/module"
)

// VoteAggregator verifies and aggregates votes to build QC.
// When enough votes have been collected, it builds a QC and send it to the EventLoop
// VoteAggregator also detects protocol violation, including invalid votes, double voting etc, and
// notifies a HotStuff consumer for slashing.
type VoteAggregator interface {
	module.ReadyDoneAware
	module.Startable

	// AddVote verifies and aggregates a vote.
	// The voting block could either be known or unknown.
	// If the voting block is unknown, the vote won't be processed until AddBlock is called with the block.
	// This method can be called concurrently, votes will be queued and processed asynchronously.
	AddVote(vote *model.Vote)

	// AddBlock notifies the VoteAggregator that it should start processing votes for the given block.
	// AddBlock is a _synchronous_ call (logic is executed by the calling go routine). It also verifies
	// validity of the proposer's vote for its own block.
	// This method can be called concurrently, block will be queued and processed asynchronously.
	AddBlock(block *model.Proposal)

	// InvalidBlock notifies the VoteAggregator about an invalid proposal, so that it
	// can process votes for the invalid block and slash the voters. Expected error
	// returns during normal operations:
	// * mempool.BelowPrunedThresholdError if proposal's view has already been pruned
	InvalidBlock(block *model.Proposal) error

	// PruneUpToView deletes all votes _below_ to the given view, as well as
	// related indices. We only retain and process whose view is equal or larger
	// than `lowestRetainedView`. If `lowestRetainedView` is smaller than the
	// previous value, the previous value is kept and the method call is a NoOp.
	PruneUpToView(view uint64)
}

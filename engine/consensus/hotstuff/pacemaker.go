package hotstuff

import (
	"github.com/dapperlabs/flow-go/engine/consensus/hotstuff/types"
)

type PaceMaker interface {
	// CurView returns the current view
	CurView() uint64

	// UpdateCurViewWithQC will check if the given QC will allow PaceMaker to fast
	// forward to the QC's view.
	// If yes, a NewViewEvent will be returned
	UpdateCurViewWithQC(qc *types.QuorumCertificate) (*types.NewViewEvent, bool)

	// UpdateCurViewWithBlock will check if the given block will allow PaceMaker to fast forward
	// to the BlockProposal's view. If yes, the PaceMaker will update it's internal value for
	// CurView and return a NewViewEvent.
	//
	// The parameter `nextPrimary` indicates to the PaceMaker whether or not this replica is the
	// primary for the NEXT view (taking curView as reference).
	// True corresponds to this replica being the next primary.
	UpdateCurViewWithBlock(block *types.BlockProposal, nextPrimary bool) (*types.NewViewEvent, bool)

	// OnTimeout is called when a timeout, which was previously created by the PaceMaker, has
	// looped through the event loop.
	// It returns a NewViewEvent if it triggers the current view to be updated.
	OnTimeout() (*types.NewViewEvent, bool)
}

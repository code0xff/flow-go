package hotstuff

import (
	"time"

	"github.com/dapperlabs/flow-go/engine/consensus/hotstuff/types"
)

type PaceMaker struct {
	curView   uint64
	highestQC *types.QuorumCertificate
	timeout   *time.Timer
	Timeouts  chan<- *types.Timeout
}

func (p *PaceMaker) CurView() uint64 {
	return p.curView
}

func (p *PaceMaker) UpdateQC(qc *types.QuorumCertificate) *types.NewViewEvent {
	panic("TODO")
}

// ToDo change to ProcessedBlockForView(view uint64)
func (p *PaceMaker) UpdateBlock(block *types.BlockProposal) *types.NewViewEvent {
	panic("TODO")
}

// ToDo change to ProcessedQcForView(view uint64)
func (p *PaceMaker) OnLocalTimeout(timeout *types.Timeout) *types.NewViewEvent {
	panic("TODO")
}

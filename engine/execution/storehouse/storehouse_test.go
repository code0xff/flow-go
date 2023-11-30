package storehouse_test

import (
	"fmt"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/storage"
	"github.com/onflow/flow-go/utils/unittest"
	"go.uber.org/atomic"
)

var unknownBlock = unittest.IdentifierFixture()
var unknownReg = makeReg("unknown", "unknown")

func makeReg(key string, value string) flow.RegisterEntry {
	return flow.RegisterEntry{
		Key: flow.RegisterID{
			Owner: "owner",
			Key:   key,
		},
		Value: []byte(value),
	}
}

type mockFinalizedReader struct {
	headerByHeight  map[uint64]*flow.Header
	lowest          uint64
	highest         uint64
	finalizedHeight *atomic.Uint64
	finalizedCalled *atomic.Int64
}

func newMockFinalizedReader(initHeight uint64, count int) (*mockFinalizedReader, map[uint64]*flow.Header, uint64) {
	root := unittest.BlockHeaderFixture(unittest.WithHeaderHeight(initHeight))
	blocks := unittest.ChainFixtureFrom(count, root)
	headerByHeight := make(map[uint64]*flow.Header, len(blocks)+1)
	headerByHeight[root.Height] = root

	for _, b := range blocks {
		headerByHeight[b.Header.Height] = b.Header
	}

	highest := blocks[len(blocks)-1].Header.Height
	return &mockFinalizedReader{
		headerByHeight:  headerByHeight,
		lowest:          initHeight,
		highest:         highest,
		finalizedHeight: atomic.NewUint64(initHeight),
		finalizedCalled: atomic.NewInt64(0),
	}, headerByHeight, highest
}

func (r *mockFinalizedReader) FinalizedBlockIDAtHeight(height uint64) (flow.Identifier, error) {
	r.finalizedCalled.Add(1)
	finalized := r.finalizedHeight.Load()
	if height > finalized {
		return flow.Identifier{}, storage.ErrNotFound
	}

	if height < r.lowest {
		return unknownBlock, nil
	}
	return r.headerByHeight[height].ID(), nil
}

func (r *mockFinalizedReader) MockFinal(height uint64) error {
	if height < r.lowest || height > r.highest {
		return fmt.Errorf("height %d is out of range [%d, %d]", height, r.lowest, r.highest)
	}

	r.finalizedHeight.Store(height)
	return nil
}

func (r *mockFinalizedReader) FinalizedCalled() int {
	return int(r.finalizedCalled.Load())
}

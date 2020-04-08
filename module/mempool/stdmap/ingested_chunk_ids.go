package stdmap

import (
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/module/mempool/model"
)

// IngestedChunkIDs represents a concurrency-safe memory pool for ingested chunk IDs.
// By ingested chunk IDs we mean those that have a verifiable chunk for them forwarded from
// Ingest engine to the Verify engine of Verification node
type IngestedChunkIDs struct {
	*Backend
}

// NewIngestedChunkIDs creates a new memory pool for chunk states.
func NewIngestedChunkIDs(limit uint) (*IngestedChunkIDs, error) {
	i := &IngestedChunkIDs{
		Backend: NewBackend(WithLimit(limit)),
	}

	return i, nil
}

// Add will add the given chunk ID to the memory pool or it will error if
// the chunk ID is already in the memory pool.
func (i *IngestedChunkIDs) Add(chunk *flow.Chunk) error {
	// wraps chunk ID around an ID entity to be stored in the mempool
	id := &model.IdEntity{
		Id: chunk.ID(),
	}
	return i.Backend.Add(id)
}

// Has checks whether the mempool has the chunk ID
func (i *IngestedChunkIDs) Has(chunkID flow.Identifier) bool {
	return i.Backend.Has(chunkID)
}

// Rem will remove the given chunk ID from the memory pool; it will
// return true if the chunk ID was known and removed.
func (i *IngestedChunkIDs) Rem(chunkID flow.Identifier) bool {
	return i.Backend.Rem(chunkID)
}

// All returns all chunk IDs stored in the mempool
func (i *IngestedChunkIDs) All() flow.IdentifierList {
	entities := i.Backend.All()
	chunkIDs := make([]flow.Identifier, 0, len(entities))
	for _, entity := range entities {
		idEntity := entity.(model.IdEntity)
		chunkIDs = append(chunkIDs, idEntity.Id)
	}
	return chunkIDs
}

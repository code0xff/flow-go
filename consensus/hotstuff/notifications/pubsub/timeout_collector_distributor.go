package pubsub

import (
	"sync"

	"github.com/onflow/flow-go/consensus/hotstuff"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/model/flow"
)

// TimeoutCollectorDistributor ingests events from hotstuff's VoteCollector
// logic and distributes them to subscribers. Concurrently safe.
// TODO: investigate if this can be updated using atomics to prevent locking on mutex since we always add all consumers
// before delivering events.
type TimeoutCollectorDistributor struct {
	lock      sync.RWMutex
	consumers []hotstuff.TimeoutCollectorConsumer
}

var _ hotstuff.TimeoutCollectorConsumer = (*TimeoutCollectorDistributor)(nil)

func NewTimeoutCollectorDistributor() *TimeoutCollectorDistributor {
	return &TimeoutCollectorDistributor{}
}

func (d *TimeoutCollectorDistributor) AddTimeoutCollectorConsumer(consumer hotstuff.TimeoutCollectorConsumer) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.consumers = append(d.consumers, consumer)
}

func (d *TimeoutCollectorDistributor) OnTcConstructedFromTimeouts(tc *flow.TimeoutCertificate) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	for _, consumer := range d.consumers {
		consumer.OnTcConstructedFromTimeouts(tc)
	}
}

func (d *TimeoutCollectorDistributor) OnPartialTcCreated(view uint64, newestQC *flow.QuorumCertificate, lastViewTC *flow.TimeoutCertificate) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	for _, consumer := range d.consumers {
		consumer.OnPartialTcCreated(view, newestQC, lastViewTC)
	}
}

func (d *TimeoutCollectorDistributor) OnNewQcDiscovered(qc *flow.QuorumCertificate) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	for _, consumer := range d.consumers {
		consumer.OnNewQcDiscovered(qc)
	}
}

func (d *TimeoutCollectorDistributor) OnNewTcDiscovered(tc *flow.TimeoutCertificate) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	for _, consumer := range d.consumers {
		consumer.OnNewTcDiscovered(tc)
	}
}

func (d *TimeoutCollectorDistributor) OnTimeoutProcessed(timeout *model.TimeoutObject) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	for _, subscriber := range d.consumers {
		subscriber.OnTimeoutProcessed(timeout)
	}
}

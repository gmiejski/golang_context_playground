package metrics

type SpeculativeExecutionMetric interface {
	Increment()
	GetCancelled() int
}

type AtomicSpeculativeExecutionMetric struct {

}

func (a *AtomicSpeculativeExecutionMetric) Increment() {
	panic("implement me")
}

func (a *AtomicSpeculativeExecutionMetric) GetCancelled() int {

	return 0
}

func NewMetrics() SpeculativeExecutionMetric {
	return &AtomicSpeculativeExecutionMetric{}
}

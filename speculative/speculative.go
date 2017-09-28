package speculative

import (
	"org.miejski/contextPlayground/metrics"
	"org.miejski/contextPlayground/service"
	"context"
	"time"
)

type SpeculativeExecutionClient interface {
	Get() (int, error)
}

type SmartSpeculativeExecutionClient struct {
	Client  service.DependencyClient
	Metrics metrics.SpeculativeExecutionMetric
	Hosts   []service.Host
}

func (client *SmartSpeculativeExecutionClient) Get() (int, error) {
	result := make(chan int)
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()
	for _, h := range client.Hosts {
		go client.asyncGet(ctx, h, result)
	}

	select {
	case resultValue, ok := <-result:
		println("Received first result ")
		if !ok {
			println("Closed")
			return -1, nil
		}
		println(resultValue)
		return 1, nil
	case <-ctx.Done():
		println("Done")
		return -1, ctx.Err()
	}
}

func (m *SmartSpeculativeExecutionClient) asyncGet(ctx context.Context, host service.Host, channel chan int) {
	thisChan := make(chan int)
	call := func() {
		value := m.Client.Get(host)
		thisChan <- value
	}
	go call()
	select {
	case <-ctx.Done():
		println("Context has ended")
	case s := <-thisChan:
		print("received")
		channel <- s
	}
}

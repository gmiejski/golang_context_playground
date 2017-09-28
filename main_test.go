package main

import (
	"testing"
	s "org.miejski/contextPlayground/service"
	m "org.miejski/contextPlayground/metrics"
	"time"
	"org.miejski/contextPlayground/speculative"
)

func BenchmarkSpeculativeExecution(b *testing.B) {

	hosts := []s.Host{s.Host(100 * time.Millisecond), s.Host(200 * time.Millisecond), s.Host(300 * time.Millisecond), s.Host(400 * time.Millisecond)}
	dependencyClient := s.NewDependencyClient()
	service := speculative.SmartSpeculativeExecutionClient{Client: dependencyClient, Metrics: m.NewMetrics(), Hosts: hosts}

	for n := 0; n < b.N; n++ {
		service.Get() // TODO go ?
	}

	speculativeExecutionsCancelled := service.Metrics.GetCancelled()
	print(speculativeExecutionsCancelled)
}

package main

import (
	"org.miejski/contextPlayground/service"
	"org.miejski/contextPlayground/metrics"
	"org.miejski/contextPlayground/speculative"
)

func main() {

	hosts := []service.Host{service.Host(100), service.Host(200), service.Host(300), service.Host(400)}

	dependencyClient := service.NewDependencyClient()
	b := speculative.SmartSpeculativeExecutionClient{Metrics: metrics.NewMetrics(), Client: dependencyClient,Hosts:hosts}
	r, err := b.Get()

	if err != nil {
		println(err.Error())
		return
	}
	println("Finished")
	println(r)
}

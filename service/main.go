package service

import (
	"time"
)

type DependencyClient interface {
	Get(host Host) int
}

type MemoryDependencyClient struct {
}

func NewDependencyClient() DependencyClient {
	return &MemoryDependencyClient{}
}

func (m *MemoryDependencyClient) Get(host Host) int {
	i := time.Duration(int(host)) * time.Millisecond
	time.Sleep(i)
	return int(host)
}

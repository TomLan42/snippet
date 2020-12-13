package main

/*
 * @Author: TomLan42
 * @Date: 2020-12-10 23:19:08
 * @Last Modified by: TomLan42
 * @Last Modified time: 2020-12-10 23:41:56
 */

import (
	"fmt"
	"math/rand"
)

// Server represents a weighted server
type Server struct {
	Name   string
	Weight int
}

// Run the workload
func (s *Server) Run() {
	fmt.Printf("Server %s fired.\n", s.Name)
}

// LoadBalancer interface
type LoadBalancer interface {
	Schedule()
	Test()
}

// BaseLoadBalancer is the base class of the lb
type BaseLoadBalancer struct {
	Cluster []Server
	Metric  map[string]int
}

// Test prints out the server worklaod frequency to compare with ideal cases
func (b *BaseLoadBalancer) Test() {
	total := 0
	for _, count := range b.Metric {
		total += count
	}
	for server, count := range b.Metric {
		fmt.Printf("Server %s :  %.2f %%\n", server, 100*float64(count)/float64(total))
	}

}

// LotteryLoadBalancer leverages random number generator
type LotteryLoadBalancer struct {
	BaseLoadBalancer
}

// Schedule method of LotteryScheduler
func (s *LotteryLoadBalancer) Schedule() {
	totalWeights := 0
	for _, server := range s.Cluster {
		totalWeights += server.Weight
	}
	lottery := rand.Intn(totalWeights)
	cover := 0
	for _, server := range s.Cluster {
		cover += server.Weight
		if cover > lottery {
			server.Run()
			s.Metric[server.Name]++
			return
		}
	}
}

// StepLoadBalancer uses stepsize method
type StepLoadBalancer struct {
	BaseLoadBalancer
}

// Schedule method of StepScheduler
func (s *StepLoadBalancer) Schedule([]Server) {
}

// NgnixLoadBalancer style of load balancing
type NgnixLoadBalancer struct {
	BaseLoadBalancer
}

// Schedule method of StepScheduler
func (n *NgnixLoadBalancer) Schedule([]Server) {
}

func main() {

	cluster := []Server{
		{Name: "A", Weight: 50},
		{Name: "B", Weight: 50},
		{Name: "C", Weight: 100},
	}

	lbA := LotteryLoadBalancer{
		BaseLoadBalancer{
			Cluster: cluster,
			Metric:  make(map[string]int),
		},
	}

	// lbB := StepLoadBalancer{
	// 	Cluster: cluster}
	for i := 0; i < 2000; i++ {
		lbA.Schedule()
		// lbB.Schedule(
	}
	lbA.Test()

}

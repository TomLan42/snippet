package main

/*
 * @Author: TomLan42
 * @Date: 2020-12-10 23:19:08
 * @Last Modified by: TomLan42
 * @Last Modified time: 2020-12-10 23:41:56
 */

import "fmt"

// Server represents a weighted server
type Server struct {
	Name   string
	Weight int
}

// Run the workload
func (s *Server) Run() {
	fmt.Printf("Server %s fired.", s.Name)
}

// LoadBalancer interface
type LoadBalancer interface {
	Schedule()
}

// LotteryLoadBalancer leverages random number generator
type LotteryLoadBalancer struct {
}

// Schedule method of LotteryScheduler
func (s *LotteryLoadBalancer) Schedule([]Server) {

}

// StepLoadBalancer uses stepsize method
type StepLoadBalancer struct {
}

// Schedule method of StepScheduler
func (s *StepLoadBalancer) Schedule([]Server) {

}

func main() {

	clusters := []Server{
		{Name: "A", Weight: 50},
		{Name: "B", Weight: 70},
		{Name: "C", Weight: 30},
		{Name: "D", Weight: 20},
		{Name: "E", Weight: 10},
	}

	lbA := LotteryLoadBalancer{}
	lbB := StepLoadBalancer{}
	for i := 0; i < 100; i++ {
		lbA.Schedule(clusters)
		lbB.Schedule(clusters)
	}

}

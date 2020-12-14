package main

/*
 * @Author: TomLan42
 * @Date: 2020-12-10 23:19:08
 * @Last Modified by: TomLan42
 * @Last Modified time: 2020-12-10 23:41:56
 */

import (
	"fmt"
	"math"
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

// LotteryLoadBalancer leverages random number generator
type LotteryLoadBalancer struct {
	BaseLoadBalancer
}

// StepLoadBalancer uses stepsize method
type StepLoadBalancer struct {
	BaseLoadBalancer
	StepTable map[Server]float64
}

// NgnixLoadBalancer style of load balancing
type NgnixLoadBalancer struct {
	BaseLoadBalancer
	PointTable map[Server]int
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

// Schedule method of LotteryScheduler
func (l *LotteryLoadBalancer) Schedule() {
	totalWeights := 0
	for _, server := range l.Cluster {
		totalWeights += server.Weight
	}
	lottery := rand.Intn(totalWeights)
	cover := 0
	for _, server := range l.Cluster {
		cover += server.Weight
		if cover > lottery {
			server.Run()
			l.Metric[server.Name]++
			return
		}
	}
}

// Schedule method of StepScheduler
func (s *StepLoadBalancer) Schedule() {
	// find min step
	minStep := math.MaxFloat64
	var scheduledServer Server
	for _, server := range s.Cluster {
		if s.StepTable[server] <= minStep {
			minStep = s.StepTable[server]
			scheduledServer = server
		}
	}
	s.StepTable[scheduledServer] += float64(1) / float64(scheduledServer.Weight)
	scheduledServer.Run()
	s.Metric[scheduledServer.Name]++

}

// Schedule method of StepScheduler
func (n *NgnixLoadBalancer) Schedule() {
	// find max point
	maxStep := math.MinInt32
	sum := 0
	var scheduledServer Server
	for _, server := range n.Cluster {
		sum += server.Weight
		n.PointTable[server] += server.Weight
		if n.PointTable[server] >= maxStep {
			maxStep = n.PointTable[server]
			scheduledServer = server
		}
	}
	//sum
	for server := range n.PointTable {
		if server.Name == scheduledServer.Name {
			n.PointTable[server] -= sum
		}
	}
	scheduledServer.Run()
	n.Metric[scheduledServer.Name]++
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

	lbB := StepLoadBalancer{
		BaseLoadBalancer: BaseLoadBalancer{
			Cluster: cluster,
			Metric:  make(map[string]int),
		},
		StepTable: make(map[Server]float64),
	}

	lbC := NgnixLoadBalancer{
		BaseLoadBalancer: BaseLoadBalancer{
			Cluster: cluster,
			Metric:  make(map[string]int),
		},
		PointTable: make(map[Server]int),
	}

	for _, server := range lbC.Cluster {
		lbC.PointTable[server] = 0
	}

	for i := 0; i < 2000; i++ {
		lbA.Schedule()
		lbB.Schedule()
		lbC.Schedule()
	}
	fmt.Println("***Lottery Balancer")
	lbA.Test()
	fmt.Println("***Step Balancer")
	lbB.Test()
	fmt.Println("***Step Balancer")
	lbC.Test()

}

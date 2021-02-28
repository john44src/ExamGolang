package scheduler

import (
	"exam/scheduler/cpu"
	"exam/scheduler/job"
	"exam/scheduler/mqms"
	"exam/scheduler/mqmsbasic"
	"exam/scheduler/mqmsrr"
	"exam/scheduler/mqmsws"
	"testing"
)

const (
	numCPUs = 4
)

func TestMQMSBasicSystem(t *testing.T) {
	job.ResetJobCounter()
	schedule := []*entry{
		j(t010, 0),    // a
		j(t003, 0),    // b
		j(t010, 0),    // c
		j(t003, 0),    // d
		j(t010, 0),    // e
		j(t005, 0),    // f
		j(t010, 0),    // g
		j(t008, 0),    // h
		j(t005, t007), // i
		j(t008, t009), // j
	}
	cpus := cpu.NewCPUs(numCPUs, 0)
	scheduler := mqmsbasic.New(cpus)
	sys := NewSystem(scheduler, cpus)
	sys.Run(schedule)
}

func TestMQMSRoundRobinSystem(t *testing.T) {
	job.ResetJobCounter()
	schedule := []*entry{
		j(t010, 0),    // a
		j(t003, 0),    // b
		j(t010, 0),    // c
		j(t003, 0),    // d
		j(t010, 0),    // e
		j(t005, 0),    // f
		j(t010, 0),    // g
		j(t008, 0),    // h
		j(t005, t007), // i
		j(t008, t009), // j
	}
	cpus := cpu.NewCPUs(numCPUs, 0)
	scheduler := mqmsrr.New(cpus, t005)
	sys := NewSystem(scheduler, cpus)
	sys.Run(schedule)
}

func TestMQMSWorkStealingSystem(t *testing.T) {
	job.ResetJobCounter()
	schedule := []*entry{
		j(t010, 0),    // a
		j(t003, 0),    // b
		j(t010, 0),    // c
		j(t003, 0),    // d
		j(t010, 0),    // e
		j(t005, 0),    // f
		j(t010, 0),    // g
		j(t008, 0),    // h
		j(t005, t007), // i
		j(t008, t009), // j
	}
	cpus := cpu.NewCPUs(numCPUs, 0)
	scheduler := mqmsws.New(cpus, t005)
	sys := NewSystem(scheduler, cpus)
	sys.Run(schedule)
}

func TestMQMSConcurrentSystem(t *testing.T) {
	job.ResetJobCounter()
	schedule := []*entry{
		j(t010, 0),    // a
		j(t003, 0),    // b
		j(t010, 0),    // c
		j(t003, 0),    // d
		j(t010, 0),    // e
		j(t005, 0),    // f
		j(t010, 0),    // g
		j(t008, 0),    // h
		j(t005, t007), // i
		j(t008, t009), // j
	}
	cpus := cpu.NewCPUs(numCPUs, 0)
	scheduler := mqms.New(cpus, t005)
	sys := NewSystem(scheduler, cpus)
	sys.Run(schedule)
}

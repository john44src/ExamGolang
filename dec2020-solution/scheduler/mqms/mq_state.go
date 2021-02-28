package mqms

import (
	"exam/dec2020-solution/scheduler/config"
	"exam/dec2020-solution/scheduler/cpu"
	"exam/dec2020-solution/scheduler/job"
)

type mqState struct {
	cpu   *cpu.CPU
	queue chan *job.Job
}

func NewMQState(cpu *cpu.CPU) *mqState {
	return &mqState{
		cpu:   cpu,
		queue: make(chan *job.Job, config.QueueSize),
	}
}

func (s *mqState) Tick() bool {
	return s.cpu.Tick()
}

func (s *mqState) Len() int {
	return len(s.queue)
}

func (s *mqState) Add(job *job.Job) {
	if job != nil {
		// only add real jobs
		s.queue <- job
	}
}

func (s *mqState) Remove() *job.Job {
	if len(s.queue) == 0 {
		return nil
	}
	removedJob := <-s.queue
	return removedJob
}

func (s *mqState) Reassign() {
	currentJob := s.cpu.CurrentJob()
	if currentJob != nil {
		// add current job to end of CPU queue
		s.Add(currentJob)
	}
	// find new job to run on this CPU
	nxtJob := s.Remove()
	s.cpu.Assign(nxtJob)
}

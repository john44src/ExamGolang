package mqmsrr

import (
	"exam/dec2020-solution/scheduler/cpu"
	"exam/dec2020-solution/scheduler/job"
)

type mqState struct {
	cpu   *cpu.CPU
	queue job.Jobs
}

func NewMQState(cpu *cpu.CPU) *mqState {
	return &mqState{
		cpu:   cpu,
		queue: make(job.Jobs, 0),
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
		s.queue = append(s.queue, job)
	}
}

func (s *mqState) Remove() *job.Job {
	if len(s.queue) == 0 {
		return nil
	}
	removedJob := s.queue[0]
	s.queue = s.queue[1:]
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

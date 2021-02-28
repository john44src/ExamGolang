package mqmsrr

import (
	"exam/dec2020-solution/scheduler/config"
	"exam/dec2020-solution/scheduler/cpu"
	"exam/dec2020-solution/scheduler/job"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	t005 = config.T005
	t010 = config.T010
	t020 = config.T020
)

var jobs = job.Jobs{
	job.New(10, t010), // a
	job.New(10, t010), // b
	job.New(10, t005), // c
	job.New(10, t020), // d
	job.New(10, t020), // e
}

func TestMQStateAdd(t *testing.T) {
	cpu := cpu.NewCPU(0, 0)
	s := NewMQState(cpu)
	s.Add(jobs[0])
	if s.Len() != 1 {
		t.Errorf("Len() = %d, expected %d", s.Len(), 1)
	}
	s.Add(jobs[0])
	if s.Len() != 2 {
		t.Errorf("Len() = %d, expected %d", s.Len(), 2)
	}
	for _, j := range jobs {
		s.Add(j)
	}
	if s.Len() != 2+len(jobs) {
		t.Errorf("Len() = %d, expected %d", s.Len(), 2+len(jobs))
	}
}

func TestMQStateAddRemove(t *testing.T) {
	cpu := cpu.NewCPU(0, 0)
	s := NewMQState(cpu)
	s.Add(jobs[0])
	if s.Len() != 1 {
		t.Errorf("Len() = %d, expected %d", s.Len(), 1)
	}
	j := s.Remove()
	if s.Len() != 0 {
		t.Errorf("Len() = %d, expected %d", s.Len(), 0)
	}
	if !cmp.Equal(jobs[0], j) {
		t.Errorf("s.add(%v); %v != s.remove(); expected removed job to be same", jobs[0], j)
	}
}

func TestMQStateAddRemoveMany(t *testing.T) {
	cpu := cpu.NewCPU(0, 0)
	s := NewMQState(cpu)
	for _, j := range jobs {
		s.Add(j)
	}
	var js job.Jobs
	for i := s.Len(); i > 0; i-- {
		j := s.Remove()
		js = append(js, j)
	}
	if s.Len() != 0 {
		t.Errorf("Len() = %d, expected %d", s.Len(), 0)
	}
	for i := range js {
		if !cmp.Equal(jobs[i], js[i]) {
			t.Errorf("jobs[%d] = %v != js[%d] = %v; expected same job at index %d", i, jobs[i], i, js[i], i)
		}
	}
}

func TestMQStateAddRemoveTooMany(t *testing.T) {
	cpu := cpu.NewCPU(0, 0)
	s := NewMQState(cpu)
	for _, j := range jobs {
		s.Add(j)
	}
	for i := s.Len(); i > 0; i-- {
		s.Remove()
	}
	if s.Len() != 0 {
		t.Errorf("Len() = %d, expected %d", s.Len(), 0)
	}
	j := s.Remove()
	if j != nil {
		t.Errorf("Remove() = %v, expected nil", j)
	}
}

func TestMQStateReassign(t *testing.T) {
	cpu := cpu.NewCPU(0, 0)
	s := NewMQState(cpu)
	for _, j := range jobs {
		s.Add(j)
	}
	currentJob := s.cpu.CurrentJob()
	if currentJob != nil {
		t.Errorf("Found job %v on CPU%d; expected no jobs", currentJob, s.cpu.ID())
	}
	s.Reassign()
	newJob := s.cpu.CurrentJob()
	if currentJob == newJob {
		t.Errorf("Reassign(): expected new job %v to be different from previous %v", newJob, currentJob)
	}
	s.Reassign()
	newJob = s.cpu.CurrentJob()
	if currentJob == newJob {
		t.Errorf("Reassign(): expected new job %v to be different from previous %v", newJob, currentJob)
	}
}

func TestMQStateReassignMany(t *testing.T) {
	cpu := cpu.NewCPU(0, 0)
	s := NewMQState(cpu)
	for _, j := range jobs {
		s.Add(j)
	}
	for _, j := range jobs {
		s.Reassign()
		newJob := s.cpu.CurrentJob()
		if j != newJob {
			t.Errorf("Reassign(): expected new job %v to be different from previous %v", newJob, j)
		}
	}
	// cycle around
	for _, j := range jobs {
		s.Reassign()
		newJob := s.cpu.CurrentJob()
		if j != newJob {
			t.Errorf("Reassign(): expected new job %v to be different from previous %v", newJob, j)
		}
	}
	lastJob := s.cpu.CurrentJob()
	if lastJob != jobs[len(jobs)-1] {
		t.Errorf("Reassign(): expected current job %v to be last job %v in list", lastJob, jobs[len(jobs)-1])
	}
	s.Reassign()
	firstJob := s.cpu.CurrentJob()
	if firstJob != jobs[0] {
		t.Errorf("Reassign(): expected new job %v to be first job %v in list", firstJob, jobs[0])
	}
}

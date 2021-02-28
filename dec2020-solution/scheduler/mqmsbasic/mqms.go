package mqmsbasic

import (
	"exam/dec2020-solution/scheduler/cpu"
	"exam/dec2020-solution/scheduler/job"
	"time"
)

// Multi-queue multiprocessor scheduler (MQMS)
type multiQueueScheduler struct {
	state []*mqState
}

func New(cpus []*cpu.CPU) *multiQueueScheduler {
	schedulerState := make([]*mqState, len(cpus))
	for i := range cpus {
		schedulerState[i] = NewMQState(cpus[i])
	}
	return &multiQueueScheduler{
		state: schedulerState,
	}
}

// Add jobs to queue and set arrival time for the job.
// The job is added to exactly one queue with the fewest jobs.
// The choice of shortest queue should be deterministic, starting from CPU0 and up.
func (q *multiQueueScheduler) Add(job *job.Job) {
	job.Scheduled()
	i := q.findShortestQueue()
	q.state[i].Add(job)
}

// findShortestQueue returns the index of the scheduling queue with the fewest jobs.
// If multiple queues have the same number of jobs, the first queue found is returned.
func (q *multiQueueScheduler) findShortestQueue() int {
	shortestQueueLen := 10000
	shortestQueueIndex := 0
	for i, s := range q.state {
		if s.Len() < shortestQueueLen {
			shortestQueueLen = s.Len()
			shortestQueueIndex = i
		}
	}
	return shortestQueueIndex
}

// Tick runs the scheduled jobs for the system time on all CPUs, and returns
// the number of jobs finished in this tick. Depending on scheduler requirements,
// the Tick method may assign new jobs to the different CPUs before returning.
func (q *multiQueueScheduler) Tick(systemTime time.Duration) int {
	jobsFinished := 0
	for _, state := range q.state {
		if state.cpu.IsRunning() {
			// job running on this CPU is not done yet
			if state.Tick() {
				// job that was running on this CPU is done
				jobsFinished++
				state.Reassign()
			}
		} else {
			// CPU is idle, try to reassign from own queue first
			state.Reassign()
		}
	}
	return jobsFinished
}

package mqms

import (
	"exam/dec2020-solution/scheduler/cpu"
	"exam/dec2020-solution/scheduler/job"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Multi-queue multiprocessor scheduler (MQMS)
type multiQueueScheduler struct {
	state   []*mqState
	quantum time.Duration
}

func New(cpus []*cpu.CPU, quantum time.Duration) *multiQueueScheduler {
	schedulerState := make([]*mqState, len(cpus))
	for i := range cpus {
		schedulerState[i] = NewMQState(cpus[i])
	}
	return &multiQueueScheduler{
		state:   schedulerState,
		quantum: quantum,
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

// findLongestQueue returns the index of the scheduling queue with the most jobs.
// If multiple queues have the same number of jobs, the first queue found is returned.
func (q *multiQueueScheduler) findLongestQueue() int {
	longestQueueLen := 0
	longestQueueIndex := 0
	for i, s := range q.state {
		if s.Len() > longestQueueLen {
			longestQueueLen = s.Len()
			longestQueueIndex = i
		}
	}
	return longestQueueIndex
}

// Tick runs the scheduled jobs for the system time on all CPUs, and returns
// the number of jobs finished in this tick. Depending on scheduler requirements,
// the Tick method may assign new jobs to the different CPUs before returning.
func (q *multiQueueScheduler) Tick(systemTime time.Duration) int {
	var jobsFinished uint64
	var wg sync.WaitGroup
	sliceExpired := systemTime%q.quantum == 0
	for _, s := range q.state {
		wg.Add(1)
		go func(state *mqState) {
			defer wg.Done()

			if state.cpu.IsRunning() {
				// job running on this CPU is not done yet
				if state.Tick() {
					// job that was running on this CPU is done
					atomic.AddUint64(&jobsFinished, 1)
				}
			}

			if sliceExpired {
				// try to reassign from own queue first
				state.Reassign()
				if !state.cpu.IsRunning() {
					// if after scheduling from own queue, CPU is found to be idle,
					// try to steal work from other queue
					i := q.findLongestQueue()
					if q.state[i].Len() > 0 {
						stolenJob := q.state[i].Remove()
						fmt.Printf("%v: CPU%d no more jobs: stealing job from CPU%d: %v\n", systemTime, state.cpu.ID(), i, stolenJob)
						state.cpu.Assign(stolenJob)
					}
				}
			}
		}(s)
	}
	wg.Wait()
	return int(jobsFinished)
}

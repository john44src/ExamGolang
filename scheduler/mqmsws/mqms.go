package mqmsws

import (
	"exam/scheduler/cpu"
	"exam/scheduler/job"
	"time"
)

// Multi-queue multiprocessor scheduler (MQMS)
type multiQueueScheduler struct {
}

func New(cpus []*cpu.CPU, quantum time.Duration) *multiQueueScheduler {
	return &multiQueueScheduler{}
}

// Add jobs to queue and set arrival time for the job.
// The job is added to exactly one queue with the fewest jobs.
// The choice of shortest queue should be deterministic, starting from CPU0 and up.
func (q *multiQueueScheduler) Add(job *job.Job) {
}

// Tick runs the scheduled jobs for the system time on all CPUs, and returns
// the number of jobs finished in this tick. Depending on scheduler requirements,
// the Tick method may assign new jobs to the different CPUs before returning.
func (q *multiQueueScheduler) Tick(systemTime time.Duration) int {
	jobsFinished := 0
	return jobsFinished
}

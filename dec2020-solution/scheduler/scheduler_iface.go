package scheduler

import (
	"exam/dec2020-solution/scheduler/job"
	"time"
)

type Scheduler interface {
	Add(*job.Job)
	Tick(time.Duration) int
}

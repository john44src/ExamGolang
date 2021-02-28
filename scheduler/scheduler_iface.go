package scheduler

import (
	"exam/scheduler/job"
	"time"
)

type Scheduler interface {
	Add(*job.Job)
	Tick(time.Duration) int
}

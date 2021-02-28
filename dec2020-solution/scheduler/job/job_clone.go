package job

import (
	"fmt"
	"time"
)

func NewTestJob(id int, estimated, remaining time.Duration) Job {
	return Job{
		id:        id,
		estimated: estimated,
		remaining: remaining,
	}
}

func (j *Job) Clone() Job {
	if j == nil {
		return *j
	}
	return Job{
		id:        j.id,
		wsSize:    j.wsSize,
		start:     j.start,
		arrival:   j.arrival,
		estimated: j.estimated,
		remaining: j.remaining,
	}
}

func (j Job) GoString() string {
	return fmt.Sprintf("jb(%d, %d, %d)", j.id, j.estimated, j.remaining)
}

func (j Job) Remaining() time.Duration {
	return j.remaining
}

// added by a student
func (j Job) EstimatedTime() (estimated time.Duration) {
	estimated = j.estimated
	return estimated
}

// added by a student
func (j Job) NewEstimatedTime(quantum time.Duration) {
	j.estimated = j.estimated - quantum
}

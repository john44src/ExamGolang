package job

import (
	"exam/dec2020-solution/scheduler/config"
	"fmt"
	"time"
)

var nextID = 0

// Job keeps track of when the job was started, its working set size,
// CPU affinities, its estimated running time and the remaining time.
type Job struct {
	id        int
	wsSize    int
	start     time.Time
	arrival   time.Time
	estimated time.Duration
	remaining time.Duration
}

// New returns a job with given working set size and estimated running time.
func New(size int, estimated time.Duration) *Job {
	nextID++
	return NewJob(nextID, size, estimated)
}

// NewJob returns a job with given working set size and estimated running time.
func NewJob(id, size int, estimated time.Duration) *Job {
	return &Job{
		id:        id,
		wsSize:    size,
		estimated: estimated,
		remaining: estimated,
	}
}

func ResetJobCounter() {
	nextID = 0
}

func (j *Job) Scheduled() {
	j.arrival = time.Now()
}

func (j *Job) Started(cpuID int) {
	j.start = time.Now()
}

// run runs the job for the given duration.
func (j *Job) run(durationToRun time.Duration) bool {
	j.remaining -= durationToRun
	return j.remaining <= 0
}

// Tick runs the job for one tick and returns true if job is finished.
func (j *Job) Tick() bool {
	return j.run(config.TickDuration)
}

// ID returns the job's ID.
func (j Job) ID() int {
	return j.id
}

// Size returns the size of the job, i.e. how much cache space it takes up.
func (j Job) Size() int {
	return j.wsSize
}

// Equal returns true if this job and the given job has the same id.
func (j Job) Equal(job Job) bool {
	return j.id == job.id && j.estimated == job.estimated && j.remaining == job.remaining
}

func (j Job) String() string {
	return fmt.Sprintf("%s(%0v)", toLetter(j.id), j.remaining)
}

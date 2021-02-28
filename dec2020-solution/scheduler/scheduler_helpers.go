package scheduler

import (
	"exam/dec2020-solution/scheduler/config"
	"exam/dec2020-solution/scheduler/job"
	"time"
)

const (
	tick = config.TickDuration
	t001 = config.T001
	t002 = config.T002
	t003 = config.T003
	t004 = config.T004
	t005 = config.T005
	t006 = config.T006
	t007 = config.T007
	t008 = config.T008
	t009 = config.T009
	t010 = config.T010
	t015 = config.T015
	t020 = config.T020
	t030 = config.T030
	t040 = config.T040
	t050 = config.T050
	t100 = config.T100
)

var jobLengths = []time.Duration{
	t001, t002, t003, t004, t005, t006, t007, t008, t009, t010, t015, t020, t030, t040, t050, t100,
}

var jobLenStrings = []string{
	"t001", "t002", "t003", "t004", "t005", "t006", "t007", "t008", "t009", "t010", "t015", "t020", "t030", "t040", "t050", "t100",
}

type entry struct {
	job     *job.Job
	arrival time.Duration
}

// j is a helper function to create a scheduling entry with a job and its arrival time.
// With this helper, the job zero size and the given estimated running time.
var j = func(estimated, arrival time.Duration) *entry {
	return &entry{job: job.New(0, estimated), arrival: arrival}
}

var ji = func(id int, estimated, arrival time.Duration) *entry {
	return &entry{job: job.NewJob(id, 0, estimated), arrival: arrival}
}

var jb = func(id int, estimated, remaining time.Duration) job.Job {
	return job.NewTestJob(id, estimated, remaining)
}

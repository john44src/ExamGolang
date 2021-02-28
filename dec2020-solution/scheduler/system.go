package scheduler

import (
	"exam/dec2020-solution/scheduler/config"
	"exam/dec2020-solution/scheduler/cpu"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type System struct {
	systemTime   time.Duration
	jobsFinished int
	scheduler    Scheduler
	cpus         []*cpu.CPU
	w            *tabwriter.Writer
}

func NewSystem(scheduler Scheduler, cpus []*cpu.CPU) *System {
	return &System{
		cpus:      cpus,
		scheduler: scheduler,
		w:         tabwriter.NewWriter(os.Stdout, 2, 8, 2, ' ', 0),
	}
}

func (s *System) Run(schedule []*entry) {
	sort.Slice(schedule, func(i, j int) bool {
		return schedule[i].arrival < schedule[j].arrival
	})
	s.printHeader()
	for s.jobsFinished < len(schedule) {
		// examine schedule for entries that just arrived, to be scheduled now
		for _, e := range schedule {
			if e.arrival == s.systemTime {
				s.scheduler.Add(e.job)
			}
		}
		s.tick()
		s.printCPUs()
	}
	s.w.Flush()
}

func (s *System) tick() {
	fmt.Fprintf(s.w, "%v\t", s.systemTime)
	s.jobsFinished += s.scheduler.Tick(s.systemTime)
	s.systemTime += config.TickDuration
}

func (s *System) printHeader() {
	fmt.Fprint(s.w, "Tick\t")
	for _, cpu := range s.cpus {
		fmt.Fprintf(s.w, "%v\t", cpu.Header())
	}
	fmt.Fprintln(s.w)
}

func (s *System) printCPUs() {
	for _, cpu := range s.cpus {
		fmt.Fprintf(s.w, "%v\t", cpu)
	}
	fmt.Fprintln(s.w)
}

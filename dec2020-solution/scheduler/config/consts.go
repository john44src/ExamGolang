package config

import "time"

const (
	// QueueSize defines the maximum number of jobs
	// that can be scheduled simultaneously.
	QueueSize = 512
)

const (
	T000         = 0
	T001         = 1 * time.Millisecond
	T002         = 2 * time.Millisecond
	T003         = 3 * time.Millisecond
	T004         = 4 * time.Millisecond
	T005         = 5 * time.Millisecond
	T006         = 6 * time.Millisecond
	T007         = 7 * time.Millisecond
	T008         = 8 * time.Millisecond
	T009         = 9 * time.Millisecond
	T010         = 10 * time.Millisecond
	T015         = 15 * time.Millisecond
	T020         = 20 * time.Millisecond
	T030         = 30 * time.Millisecond
	T040         = 40 * time.Millisecond
	T050         = 50 * time.Millisecond
	T100         = 100 * time.Millisecond
	TickDuration = T001
)

const (
	A = 1
	B = 2
	C = 3
	D = 4
	E = 5
	F = 6
)

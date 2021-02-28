package job

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

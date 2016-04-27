package job

//Job job info
type Job interface {
	Do([]byte) error
}

var jobs = make(map[string]Job)

//Register register job
func Register(n string, j Job) {
	jobs[n] = j
}

package job

import "time"

//Worker worker
type Worker interface {
	Push(string, []byte) error
	Run(time.Duration, int)
}

package cache

import "time"

//Provider cache provider
type Provider interface {
	Set(string, interface{}, time.Duration) error
	Get(string, interface{}) error
	Del(string) error
	Status() (map[string]int, error)
	Clear() error
}

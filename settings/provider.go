package settings

//Provider interface of storage
type Provider interface {
	Set(string, interface{}, bool) error
	Get(string, interface{}) error
}

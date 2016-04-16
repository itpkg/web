package dict

import (
	"os/exec"
	"strconv"
	"strings"
)

//Provider dict model
type Provider interface {
	List() (map[string]int, error)
	Query(string) (string, error)
}

//StarDict sdcv provider
type StarDict struct {
	Dir string
}

//List list all dictionaries
func (p *StarDict) List() (map[string]int, error) {
	buf, err := p.do("--list-dicts")
	if err != nil {
		return nil, err
	}
	ds := strings.Split(string(buf), "\n")
	val := make(map[string]int)
	for _, line := range ds[1:] {
		idx := strings.LastIndex(line, " ")
		if idx == -1 {
			continue
		}
		cnt, err := strconv.Atoi(line[idx+1:])
		if err != nil {
			return nil, err
		}
		val[strings.TrimSpace(line[0:idx])] = cnt
	}
	return val, nil
}

//Query search in dictionaries
func (p *StarDict) Query(key string) (string, error) {
	buf, err := p.do(key)
	return string(buf), err
}

func (p *StarDict) do(args ...string) ([]byte, error) {
	args = append([]string{"--data-dir", p.Dir}, args...)
	return exec.Command("sdcv", args...).Output()
}

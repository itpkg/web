package web_test

import (
	"testing"
	"time"

	"github.com/itpkg/web"
)

type S struct {
	K1 string
	K2 int
	K3 time.Time
}

func TestConfig(t *testing.T) {
	fn := "test.toml"
	s1 := S{K1: "hello", K2: 123, K3: time.Now()}
	if err := web.Store(fn, &s1); err != nil {
		t.Fatal(err)
	}
	var s2 S
	if err := web.Load(fn, &s2); err != nil {
		t.Fatal(err)
	}
	if s2.K2 != s1.K2 || s2.K1 != s1.K1 {
		t.Errorf("want %v+, get %v+", s1, s2)
	}
}

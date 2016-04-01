package cache_test

import (
	"testing"
	"time"

	"github.com/itpkg/web/cache"
	"github.com/itpkg/web/engines/base"
)

type S struct {
	Val int
}

const key = "kkk"

func testProvider(t *testing.T, cp cache.Provider) {
	s := S{Val: 111}
	if err := cp.Set(key, &s, 60*time.Hour); err != nil {
		t.Errorf("Bad in set: %v", err)
	}
	var s1 S
	if err := cp.Get(key, &s1); err != nil {
		t.Errorf("Bad in get: %v", err)
	}
	if s.Val != s1.Val {
		t.Errorf("Wang %d, get %d", s.Val, s1.Val)
	}
}

func TestRedis(t *testing.T) {
	re := base.Redis{Host: "localhost", Port: 6379}
	testProvider(t, &cache.RedisProvider{Redis: re.Open()})

}

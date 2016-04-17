package token_test

import (
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/itpkg/web/token"
)

func TestRedis(t *testing.T) {
	rep := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	var tp token.Provider
	tp = &token.RedisProvider{Redis: rep}

	val1 := map[string]interface{}{
		"iv": 12345,
		"tv": time.Now(),
		"fv": 1.2,
		"sv": "hello",
	}

	jwt := token.Jwt{Provider: tp, Key: []byte("hello")}

	tk, err := jwt.New(val1, 10*time.Hour*24*7)
	if err == nil {
		t.Logf("Token[%v]: %s", val1, tk)
	} else {
		t.Errorf("bad in generate token: %v", err)
	}

	if val2, err := jwt.Parse(tk); err == nil {
		t.Logf("Parse result: %v", val2)
	} else {
		t.Errorf("bad in parse token: %v", err)
	}

}

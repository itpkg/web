package token

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

//RedisProvider redis token provider
type RedisProvider struct {
	Redis *redis.Pool
}

//Set set key
func (p *RedisProvider) Set(kid string, key []byte, exp time.Duration) error {
	rc := p.Redis.Get()
	defer rc.Close()
	_, err := rc.Do("SET", p.key(kid), key, "EX", int(exp/time.Second))
	return err
}

//Get get key
func (p *RedisProvider) Get(kid string) ([]byte, error) {
	rc := p.Redis.Get()
	defer rc.Close()
	key, err := redis.Bytes(rc.Do("GET", p.key(kid)))
	if err == nil {
		return key, nil
	}
	return nil, err

}

//All list kids
func (p *RedisProvider) All() (map[string]int, error) {
	rc := p.Redis.Get()
	defer rc.Close()
	rs := make(map[string]int)
	keys, err := redis.Strings(rc.Do("KEYS", p.key("*")))
	if err != nil {
		return nil, err
	}
	for _, k := range keys {
		ttl, err := redis.Int(rc.Do("TTL", k))
		if err != nil {
			return nil, err
		}
		rs[k[len(p.key("")):]] = ttl
	}

	return rs, nil
}

//Clear clear kids
func (p *RedisProvider) Clear() error {
	rc := p.Redis.Get()
	defer rc.Close()
	keys, err := redis.Strings(rc.Do("KEYS", p.key("*")))
	if err != nil {
		return err
	}
	for _, k := range keys {
		_, err := rc.Do("DEL", k)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *RedisProvider) key(kid string) string {
	return fmt.Sprintf("token://%s", kid)
}

package cache

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"gopkg.in/vmihailenco/msgpack.v2"
)

//RedisProvider cache redis provider
type RedisProvider struct {
	Redis  *redis.Pool
	Prefix string
}

//Set set cache
func (p *RedisProvider) Set(k string, v interface{}, t time.Duration) error {

	buf, err := msgpack.Marshal(v)
	if err != nil {
		return err
	}
	rc := p.Redis.Get()
	defer rc.Close()
	_, err = rc.Do("SET", p.key(k), buf, "EX", int(t/time.Second))
	return err

}

//Get get cache
func (p *RedisProvider) Get(k string, v interface{}) error {
	rc := p.Redis.Get()
	defer rc.Close()
	buf, err := redis.Bytes(rc.Do("GET", p.key(k)))
	if err != nil {
		return err
	}
	return msgpack.Unmarshal(buf, v)

}

//Del del cache
func (p *RedisProvider) Del(k string) error {

	rc := p.Redis.Get()
	defer rc.Close()
	_, err := rc.Do("DEL", p.key(k))
	return err

}

//Status status
func (p *RedisProvider) Status() (map[string]int, error) {
	rc := p.Redis.Get()
	defer rc.Close()

	status := make(map[string]int)

	keys, err := redis.Strings(rc.Do("KEYS", p.key("*")))
	if err != nil {
		return nil, err
	}

	idx := len(p.key(""))
	for _, k := range keys {
		if ttl, err := redis.Int(rc.Do("TTL", k)); err == nil {
			status[k[idx:]] = ttl
		} else {
			return nil, err
		}
	}
	return status, nil

}

//Clear clear all caches
func (p *RedisProvider) Clear() error {

	rc := p.Redis.Get()
	defer rc.Close()

	keys, err := redis.Strings(rc.Do("KEYS", p.key("*")))
	if err != nil {
		return err
	}
	for _, k := range keys {
		if _, err = rc.Do("DEL", k); err != nil {
			return err
		}
	}
	return nil
}

func (p *RedisProvider) key(k string) string {
	return fmt.Sprintf("cache://%s", k)
}

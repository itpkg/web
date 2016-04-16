package token

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//RedisProvider redis token provider
type RedisProvider struct {
	Redis *redis.Pool
}

//Set set key
func (p *RedisProvider) Set(kid string, key []byte) error {
	rc := p.Redis.Get()
	defer rc.Close()
	_, err := rc.Do("SET", p.key(kid), key)
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

func (p *RedisProvider) key(kid string) string {
	return fmt.Sprintf("token://%s", kid)
}

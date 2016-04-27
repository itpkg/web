package job

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

//RedisWorker reids worker
type RedisWorker struct {
	Redis  *redis.Pool
	Logger *log.Logger
}

//Push add task to queue
func (p *RedisWorker) Push(qu string, ag []byte) error {
	p.Logger.Printf("add job to %s", qu)
	c := p.Redis.Get()
	defer c.Close()
	_, e := c.Do("LPUSH", p.queue(qu), ag)
	return e
}

//Run main loop
func (p *RedisWorker) Run(timeout time.Duration, threads int) {
	p.Logger.Println("workers start...")
	var queues []string
	for k := range jobs {
		p.Logger.Printf("get job %s", k)
		queues = append(queues, p.queue(k))
	}

	for i := 0; i < threads; i++ {
		go p.do(p.Redis.Get(), queues, int(timeout/time.Second))
	}

	return
}

func (p *RedisWorker) do(c redis.Conn, q []string, t int) {
	args := append(q, strconv.Itoa(t))
	val, err := redis.Values(c.Do("BRPOP", args))
	if err != nil {
		p.Logger.Print(err)
	}
	if len(val) != 2 {
		p.Logger.Print("bad pop")
	}
	n := val[0].(string)[len(p.queue("")):]
	p.Logger.Printf("get job from [%s]", n)
	f := jobs[n]
	if err := f.Do(val[1].([]byte)); err == nil {
		p.Logger.Printf("[%s] - done", n)
	} else {
		p.Logger.Printf("[%s] - %v", n, err)
	}
}

func (p *RedisWorker) queue(q string) string {
	return fmt.Sprintf("task://%s", q)
}

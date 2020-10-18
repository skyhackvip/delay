package redisdelay

import (
	"log"
	"testing"
	"time"
)

func TestRedisDelay(t *testing.T) {
	delay, err := New(1*time.Second, "test", func(data interface{}) bool {
		log.Println("do task ", data)
		return true
	})
	if err != nil {
		t.Error(err)
	}
	log.Println("start ticker...")
	delay.Start()

	task1 := Task{Id: "1", Data: "task1", Delay: 5 * time.Second}
	task2 := Task{Id: "2", Data: "task2", Delay: 8 * time.Second}
	delay.AddTask(&task1)
	delay.AddTask(&task2)
	time.Sleep(10 * time.Second)
}

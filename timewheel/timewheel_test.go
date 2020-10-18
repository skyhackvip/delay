package timewheel

import (
	"log"
	"testing"
	"time"
)

func TestTimeWheel(t *testing.T) {
	tw, err := New(1*time.Second, 10, func(data interface{}) {
		log.Println("do task", data)
	})
	if err != nil {
		t.Error(err)
	}
	log.Println("start timewheel...")
	tw.Start()
	task := Task{Id: 1, Data: "test1", Delay: 12 * time.Second}
	tw.AddTask(&task)
	time.Sleep(20 * time.Second)
}

package timeticker

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	for {
		select {
		case time := <-ticker.C:
			fmt.Println(time)
		case <-done:
			return
		}
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}

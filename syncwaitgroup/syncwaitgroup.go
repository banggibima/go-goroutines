package syncwaitgroup

import (
	"fmt"
	"sync"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello, World")

	time.Sleep(1 * time.Second)
}

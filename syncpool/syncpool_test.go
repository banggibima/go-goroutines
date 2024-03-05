package syncpool

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			return new(string)
		},
	}
	group := &sync.WaitGroup{}

	s1 := "Banggi"
	s2 := "Bima"
	s3 := "Edriantino"

	pool.Put(&s1)
	pool.Put(&s2)
	pool.Put(&s3)

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			strPtr, ok := data.(*string)
			if !ok {
				t.Error("Failed to assert type to *string")
				return
			}

			fmt.Println(*strPtr)
			pool.Put(strPtr)
		}()
	}

	group.Wait()
	fmt.Println("Complete")
}

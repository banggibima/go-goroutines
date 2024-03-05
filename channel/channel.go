package channel

import (
	"fmt"
	"time"
)

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Banggi Bima Edriantino"
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Banggi Bima Edriantino"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

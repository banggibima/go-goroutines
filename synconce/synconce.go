package synconce

var counter int = 0

func OnlyOnce() {
	counter++
}

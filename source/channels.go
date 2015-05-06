package source
import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func ping(pings chan <- string, msg string) {
	pings <- msg

}

func pong(pings <-chan string, pongs chan <- string) {
	msg := <-pings
	pongs <- msg
}

func Channels() {

	message := make(chan string)

	go func() { message <- "ping" }()

	msg := <-message
	fmt.Println(msg)

	message=make(chan string, 2)
	message <- "buffered"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)

	done := make(chan bool, 1)
	go worker(done)
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

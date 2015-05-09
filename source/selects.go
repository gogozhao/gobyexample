package source
import (
	"time"
	"fmt"
)

func selects() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second*1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second*2)
		c2 <- "two"
	}()

	for i := 0; i<2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		}
	}
}

func timeouts() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second*2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second*1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second*2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second*3):
		fmt.Println("timeout 2")
	}
}

func noblocks() {
	message := make(chan string, 1)
	signals := make(chan bool, 1)


	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}

func closing() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			if j, more := <-jobs; more {
				fmt.Println("received job:", j)
			}else {
				fmt.Println("received all jobs")
				done <- true
				break
			}
		}
	}()

	for i := 0; i<3; i++ {
		jobs <- i
	}
	close(jobs)
	<-done
}

func Selects() {
	selects()
	timeouts()
	noblocks()
	closing()
}
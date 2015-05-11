package source
import (
	"time"
	"fmt"
)


func RateLimitings() {

	requests := make(chan int, 5)

	for i := 0; i<5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond*200)
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstLimiter := make(chan time.Time, 100)

	for i := 0; i<3; i++ {
		burstLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond*200) {
			burstLimiter <- t
			fmt.Println("tick write", time.Now())
		}
	}()

	burstRequests := make(chan int, 5)
	for i := 0; i<5; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	for req := range burstRequests {
		<-burstLimiter
		fmt.Println("request", req, time.Now())
	}

	go func() {
		for {
			<-burstLimiter
			fmt.Println("tick read", time.Now())
		}
	}()

	time.Sleep(time.Second*5)
}
package source
import (
	"fmt"
	"time"
)


func workers(id int, job <-chan int, result chan <- int) {
	for j := range job {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		result <- j*2
	}
}

func Workerpools() {
	job := make(chan int, 100)
	result := make(chan int, 100)

	for i := 0; i<3; i++ {
		go workers(i, job, result)
	}

	for i := 0; i<9; i++ {
		job <- i
	}
	close(job)

	for i := 0; i<9; i++ {
		<-result
	}

}
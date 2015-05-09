package source
import (
	"time"
	"fmt"
)


func Timers() {
	time1 := time.NewTimer(time.Second*2)
	<-time1.C
	fmt.Println("Time 1 is expired")

	time2 := time.NewTimer(time.Second*2)
	go func() {
		<-time2.C
		fmt.Println("Time 2 is expired")
	}()
	time2.Stop()
	fmt.Println("Time 2 stopped")
}
package source
import (
	"time"
	"fmt"
)

func Tickers() {
	tickers := time.NewTicker(time.Millisecond*500)
	go func() {
		for t := range tickers.C {
			fmt.Println("Tick at", t)
		}

	}()

	time.Sleep(time.Millisecond*1500)
	tickers.Stop()
	fmt.Println("Tick stoped")

}

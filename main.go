package main
import (
	"fmt"
	"github.com/gogozhao/gobyexample/source"
)

func main() {

	source.HelloWorld()
	fmt.Println("")

	source.Values()
	fmt.Println("")

	source.Variables()
	fmt.Println("")

	source.Constants()
	fmt.Println("")

	source.Fors()
	fmt.Println("")

	source.Ifelse()
	fmt.Println("")

	source.Swiches()
	fmt.Println("")

	source.Arrays()
	fmt.Println("")

	source.Slices()
	fmt.Println("")

	source.Maps()
	fmt.Println("")

	source.Ranges()
	fmt.Println("")

	res := source.Plus(1, 2)
	fmt.Println("1+2 =", res)
	res = source.PlusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
	fmt.Println("")

	a, b := source.Vals()
	fmt.Println(a)
	fmt.Println(b)
	_, c := source.Vals()
	fmt.Println(c)
	fmt.Println("")

	source.Sum(1, 2)
	source.Sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	source.Sum(nums...)
	fmt.Println("")

	nextInt := source.IntSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := source.IntSeq();
	fmt.Println(newInts())
	fmt.Println("")

	fmt.Println("fact", source.Fact(7))
	fmt.Println("")

	source.Pointer()
	fmt.Println("")

	source.Structs()
	fmt.Println("")

	source.Interfaces()
	fmt.Println("")

	source.Errors()
	fmt.Println("")

	source.Goroutines()
	fmt.Println("")

	source.Channels()
	fmt.Println("")

	source.Selects()
	fmt.Println("")

	source.Timers()
	fmt.Println("")

	source.Tickers()
	fmt.Println("")

	source.Workerpools()
	fmt.Println("")

	source.RateLimitings()
	fmt.Println("")

	source.Mutexes()
	fmt.Println("")

	source.Sortings()
	fmt.Println("")

	source.Panics()
	fmt.Println("")

}
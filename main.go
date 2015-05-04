package main
import (
	"fmt"
	"github.com/gogozhao/gobyexample/source"
)

func main() {
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

}
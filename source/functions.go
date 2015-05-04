package source
import "fmt"

func Plus(a int, b int) int {
	return a+b
}

func PlusPlus(a, b, c int) int {
	return a+b+c
}

func Vals() (int, int) {
	return 3, 7
}

func Sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func IntSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}

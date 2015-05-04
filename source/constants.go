package source
import (
	"fmt"
	"math"
)

const s string = "constant"

func Constants() {

	fmt.Println(s)

	const n = 500000000

	const d = 3e20/n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	//	cannot use m (type int) as type float64 in argument to math.Sin
	//	m := n
	//	fmt.Println(math.Sin(m))
}
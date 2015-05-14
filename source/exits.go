package source
import (
	"fmt"
	"os"
)

func Exits() {

	fmt.Println("start")
	defer fmt.Println("defer stop")

	os.Exit(3)
}

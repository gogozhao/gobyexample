package source
import (
	"os"
	"fmt"
)


func Enves() {
	os.Setenv("FOO", "foo")

	fmt.Println(os.Getenv("FOO"))

	for _, v := range os.Environ() {
		fmt.Println(v)
	}
}
package source
import "os"

func Panics() {
	//	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err!=nil {
		panic(err)
	}
}
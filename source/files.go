package source
import (
	"io/ioutil"
	"fmt"
	"os"
	"io"
	"bufio"
)

func checkError(e error) {
	if e!=nil {
		panic(e)
	}
}

func read() {

	b, e := ioutil.ReadFile("/tmp/defer.txt")
	checkError(e)
	fmt.Println(string(b))

	f, e := os.Open("/tmp/defer.txt")
	checkError(e)
	defer f.Close()

	b1 := make([]byte, 5)
	n1, e := f.Read(b1)
	checkError(e)
	fmt.Println(n1, string(b1))

	o2, e := f.Seek(6, 0)
	checkError(e)
	b2 := make([]byte, 2)
	n2, e := f.Read(b2)
	fmt.Println(o2, n2, string(b2))

	o3, e := f.Seek(0, 0)
	checkError(e)
	b3 := make([]byte, 5)
	n3, e := io.ReadAtLeast(f, b3, 1)
	checkError(e)
	fmt.Println(o3, n3, string(b3))

	_, e = f.Seek(0, 0)
	checkError(e)

	r4 := bufio.NewReader(f)
	b4, e := r4.Peek(5)
	fmt.Println(string(b4))

}

func write() {

	b1 := []byte("hello\ngo\n")
	ioutil.WriteFile("/tmp/data1", b1, 0644)

	f, e := os.Create("/tmp/data2")
	checkError(e)

	b2 := []byte("go")
	o2, e := f.Write(b2)
	checkError(e)
	fmt.Println(o2)

	o3, e := f.WriteString("writes")
	checkError(e)
	fmt.Println(o3)

	f.Sync()

	w := bufio.NewWriter(f)
	checkError(e)
	w.WriteString("buffer")
	w.Flush()
}

func Files() {
	read()
	write()
}

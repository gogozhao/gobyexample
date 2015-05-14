package source
import (
	"os/exec"
	"fmt"
	"io/ioutil"
	"syscall"
	"os"
)

func Processes() {

	cmd := exec.Command("date")
	output, e := cmd.Output()
	checkError(e)
	fmt.Println(string(output))

	cmd = exec.Command("grep", "hello")
	inPipe, e := cmd.StdinPipe()
	checkError(e)
	outPipe, e := cmd.StdoutPipe()
	checkError(e)
	cmd.Start()
	inPipe.Write([]byte("hello grep\ngoodbye grep"))
	inPipe.Close()
	out, e := ioutil.ReadAll(outPipe)
	checkError(e)
	cmd.Wait()
	fmt.Println(string(out))

	cmd=exec.Command("bash", "-c", "ls -a -l -h")
	output, e=cmd.Output()
	fmt.Println(string(output))

	binary, e := exec.LookPath("ls")
	checkError(e)
	fmt.Println(binary)

	args := []string{"ls", "-lah"}

	e = syscall.Exec(binary, args, os.Environ())
	checkError(e)

}

package source
import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"flag"
)

func commands() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := strings.ToUpper(scanner.Text())
		fmt.Println(s)
	}

	if e := scanner.Err(); e!=nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func commandArgs() {
	fmt.Println(os.Args)
}

func commandFlag() {
	wordPtr := flag.String("world", "foo", "a string")
	numbPtr := flag.Int("numb", 43, "a int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:",)
}

func Commands() {
}

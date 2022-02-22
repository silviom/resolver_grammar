//go:generate goyacc -o resolver.go -p "resolver" -v resolver.output resolver.y

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// resolverDebug = 4
	resolverErrorVerbose = true
	interpreter := interpreter{}

	for {
		fmt.Println("ready ...")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Bye.")
			return
		}

		input = strings.TrimSuffix(input, "\n")

		fmt.Printf(">%s<\n", input)

		if !interpreter.evaluationFailed {
			fmt.Printf("::%s\n", interpreter.eval(interpreter.parseResult))
		}
	}
}

type interpreter struct {
	input            string
	evaluationFailed bool
	parseResult      expr
}

func (i *interpreter) Error(e string) {
	fmt.Println(e)
	i.evaluationFailed = true
}

func (i *interpreter) Parse(input string) expr {
	i.input = input
	i.evaluationFailed = false
	resolverParse(i)
	return i.parseResult
}

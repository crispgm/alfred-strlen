package main

import (
	"fmt"
	"os"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

const (
	lenByRune = "rune"
	lenBySize = "size"
)

func lengthOf(input, flag string) int {
	var result int
	if flag == lenByRune {
		result = len([]rune(input))
	} else {
		// baiscally len by size
		result = len(input)
	}

	return result
}

func run() {
	var (
		flag   string
		input  string
		result int
	)
	argc := len(os.Args)
	if argc < 3 {
		return
	}
	flag = os.Args[1]
	input = os.Args[2]
	result = lengthOf(input, flag)
	wf.NewItem(input).Arg(fmt.Sprintf("%d", result)).Subtitle(fmt.Sprintf("%d", result)).Valid(true)
	wf.SendFeedback()
}

func main() {
	wf = aw.New()
	wf.Run(run)
}

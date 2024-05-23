package main

import (
	"fmt"
	"os"
	"strings"

	"fs/funcs"
)

func main() {
	arguments := os.Args

	err := funcs.ErrArgs(len(arguments))
	if err != nil {
		fmt.Println("Usage: ", err)
		return
	}

	var text string
	var style string
	var FileName string
	if len(arguments) == 2 {
		style = "standard.txt"
		text = arguments[1]
	} else if len(arguments) == 3 {
		if strings.HasPrefix(arguments[1], "--output=") {
			FileName = arguments[1][9:]
			text = arguments[2]
			style = "standard.txt"
		} else {
			text = arguments[1]
			style = arguments[2]
		}
	} else if len(arguments) == 4 {
		if !strings.HasPrefix(arguments[1], "--output=") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		} else {
			if arguments[1] == "--output=" {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
				return
			}
			FileName = arguments[1][9:]
			text = arguments[2]
			style = arguments[3]
		}
	}
	if len(style) > 0 {
		if !strings.HasSuffix(style, ".txt") {
			style += ".txt"
		}
	}

	if len(FileName) > 0 {
		if !strings.HasSuffix(FileName, ".txt") {
			FileName += ".txt"
		}
	}

	if FileName == "standard.txt" || FileName == "shadow.txt" || FileName == "thinkertoy.txt" {
		fmt.Println("Don't use this names <standard - shadow - thinkertoy> on your output file")
		return
	}

	err = funcs.ErrInput(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(FileName) > 0 {
		Output := funcs.Print(text, style)
		funcs.Output(FileName, Output)
	} else {
		fmt.Printf("%s", funcs.Print(text, style))
	}
}

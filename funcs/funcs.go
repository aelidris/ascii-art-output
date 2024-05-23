package funcs

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func IsEmpty(str string) bool {
	return len(str) == 0
}

func ErrInput(input string) error {
	for _, ch := range input {
		if ch < 32 || ch > 126 {
			return fmt.Errorf("please use just the characters from 32 to 126 in ASCII code")
		}
	}
	return nil
}

func ErrArgs(num int) error {
	if num < 2 || num > 4 {
		return errors.New("go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}
	return nil
}

// for handling the case where the input content contains only newline characters
func IsFullNewLine(slice []string) bool {
	for _, ch := range slice {
		if ch != "" {
			return false
		}
	}
	return true
}

func Print(input string, style string) string {
	databyte, err := os.ReadFile(style)
	if err != nil {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}

	content := string(databyte)
	content = strings.ReplaceAll(content, "\r", "") // for handle the error into the thinkertoy file
	letter := strings.Split(content[1:], "\n\n")    // split characters in the source file
	res := ""
	if len(input) == 0 {
		return ""
	}

	slice := strings.Split(input, "\\n") // split the input by new line for print slice by slice
	for g := 0; g < len(slice); g++ {
		if IsEmpty(slice[g]) { // if the index is empty return new line
			res += "\n"
			if g == len(slice)-1 && IsFullNewLine(slice) {
				res = res[:len(res)-1]
			}
		} else {
			for i := 0; i < 8; i++ { // print word by word and line by line
				for k := 0; k < len(slice[g]); k++ { // print line by line of the words
					word := letter[byte(slice[g][k])-32] // target the character in the source file
					row := strings.Split(word, "\n")

					for j := 0; j < len(row[i]); j++ { // print line by line

						if row[i][j] != '\n' {
							res += string(row[i][j])
						} else {
							break // stop on the last line of each character
						}
					}
				}
				res += "\n"
			}
		}
	}
	return res
}

func Output(FileName, Output string) {
	OutputFile, err := os.Create(FileName)
	if err != nil {
		log.Fatal("Unable to create the output file")
	}

	defer OutputFile.Close()

	writer := bufio.NewWriter(OutputFile)
	_, err = writer.WriteString(Output)
	if err != nil {
		fmt.Println("Error occurred while writing the output file")
	}
	err = writer.Flush()
	if err != nil {
		return
	}
}

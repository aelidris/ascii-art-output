package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestFunc(t *testing.T) {
	inputs := []string{`\n`, ``, `hello`, `hello\n`, `hello\nthis`, `\nhello`, `last try`, `hello\n\nthere`}
	styles := []string{`shadow_compare.txt`, `standard_compare.txt`, `thinkertoy_compare.txt`}
	style := []string{"shadow.txt", `standard.txt`, `thinkertoy.txt`}
	for i := 0; i < 3; i++ {
		for j := 0; j < len(inputs); j++ {

			databyte, err := os.ReadFile("compare/" + styles[i])
			if err != nil {
				log.Fatal()
			}
			content := string(databyte)
			slice := strings.Split(content, "t\n")
			cmd := exec.Command("go", "run", "./main.go", "--output=compare/test", inputs[j], style[i])
			_, err = cmd.Output()
			if err != nil {
				t.Errorf("Error executing command: %v   %d", err, j)
				return
			}

			testFile, err := os.ReadFile("compare/test.txt")
			if err != nil {
				log.Fatal()
			}
			output := string(testFile)

			if i == 0 && j == 0 {
				fmt.Println("****************** shadow test ******************")
			} else if i == 1 && j == 0 {
				fmt.Println("\n\n****************** standard test ******************")
			} else if i == 2 && j == 0 {
				fmt.Println("\n\n****************** thinkertoy test ******************")
			}
			fmt.Println(output)
			fmt.Println("-------------------------------")
			fmt.Println(slice[j])
			fmt.Println("******************************\n******************************")
			file := ""
			if i == 0 {
				file = "shadow"
			} else if i == 1 {
				file = "standard"
			} else {
				file = "thinkertoy"
			}
			if output != slice[j] {
				t.Errorf("%v in %v: wrong output", j+1, file)
			}
		}
	}
}

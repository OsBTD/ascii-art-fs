package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ContainsOnly(char string) bool {
	for i := 0; i < len(char); i++ {
		if !strings.ContainsAny(string(char[i]), "\\n") {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || len(args) > 2 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}
	input := args[0]
	if !strings.HasSuffix(args[1], ".txt") {
		args[1] += ".txt"
	}
	if args[1] != "standard.txt" && args[1] != "thinkertoy.txt" && args[1] != "shadow.txt" {
		fmt.Println("Error : this style is unavailable \nPlease choose one of the available styles \n1 : standard \n2 : thinkertoy \n3 : shadow")
		return
	}
	Banner := "standard.txt"
	if len(args) == 2 {
		Banner = args[1]
	}

	var content []byte
	var err error
	if ContainsOnly(input) {
		for i := 0; i < len(input)/2; i++ {
			fmt.Println()
		}
		return

	}

	inputsplit := strings.Split(input, "\\n")
	for _, line := range inputsplit {
		for _, c := range line {
			if c < 32 || c > 126 {
				log.Fatal("Error : input should only contain printable ascii characters")
			}
		}
	}
	content, err = os.ReadFile(Banner)
	if err != nil {
		log.Fatal("Error : couldn't read file")
	}

	Replace := make(map[rune][]string)
	Char := 32
	noreturn := strings.ReplaceAll(string(content), "\r", "")
	Lines := strings.Split(noreturn, "\n")
	for i := 0; i < len(Lines); i += 9 {
		if i+9 <= len(Lines)-1 {
			Replace[rune(Char)] = Lines[i+1 : i+9]
		}
		if Char <= 126 {
			Char++
		}

	}

	for _, line := range inputsplit {
		if line == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for j := 0; j < len(line); j++ {
				inputrune := rune(line[j])
				fmt.Print(Replace[inputrune][i])
			}
			fmt.Println()
		}
	}

}

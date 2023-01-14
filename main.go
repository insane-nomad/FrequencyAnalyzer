package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type data struct {
	word  string
	count int
}

func readFile(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}

func analyze(text string) {
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\r", " ")
	regexp, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	text = regexp.ReplaceAllString(text, " ")
	//split := strings.Split(text, " ")
	split := strings.Fields(text)
	fmt.Println(split)
}

func main() {

	text := readFile("text.txt")
	analyze(text)

	fmt.Println(text)
}

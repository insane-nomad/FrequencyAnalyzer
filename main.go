package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	bytes, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	text := string(bytes)
	replacer := strings.NewReplacer("’", "", ",", "", ".", "", "”", "", "“", "", "?", "", "!", "")
	text = replacer.Replace(text)
	fmt.Println(text)
}

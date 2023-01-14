package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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

func analyze(text string) []data {
	karta := make(map[string]int)
	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\r", " ")
	regexpa, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	text = regexpa.ReplaceAllString(text, " ")
	split := strings.Fields(text)

	for _, val := range split {
		if len(val) > 3 {
			karta[val]++
		}
	}
	sliceData := make([]data, 0, len(karta)/2)
	for key, val := range karta {
		switch val {
		case 1:
			delete(karta, key) // удаляем ключи, значения которых ==1, чтобы меньше аппендить в слайс
		default:
			sliceData = append(sliceData, data{key, val})
		}
	}
	sort.Slice(sliceData, func(i, j int) bool {
		return sliceData[i].count > sliceData[j].count
	})
	return sliceData
}

func main() {
	text := readFile("text.txt")
	result := analyze(text)

	for i := 0; i < 10; i++ {
		fmt.Printf("Слово \"%v\" встречается в тексте раз: %v\n", result[i].word, result[i].count)
	}

}

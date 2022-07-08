package main

import (
	"fmt"
	"strings"
)

func wordfrequency(str string) map[string]int {
	sentence := strings.Fields(str)
	m := make(map[string]int)

	for _, word := range sentence {
		m[word] = m[word] + 1
	}

	return m

}
func main() {
	sentence := "The day at the beach was fun fun fun"
	for key, value := range wordfrequency(sentence) {
		fmt.Println(key, "--", value)
		max := 0
		if value > max {
			max = value
			fmt.Println(key, "word has highest frequency of", max)
		}
	}
}

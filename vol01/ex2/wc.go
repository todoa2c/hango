package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	wcMap := make(map[string]int)
	for scanner.Scan() {
		text := scanner.Text()
		words := strings.Split(text, " ")
		for _, word := range words {
			wcMap[strings.TrimSpace(word)]++
		}
	}

	for w, c := range wcMap {
		fmt.Println(w, " : ", c)
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"main/datastructures/hashtable"
	"main/datastructures/heap"
	"os"
	"regexp"
	"slices"
	"strings"
)

type WordCount struct {
	Word  string
	Count int
}

var (
	wordPattern = regexp.MustCompile(`[a-zA-Z]+`)
	// wordPattern   = regexp.MustCompile(`[^\s.,();{}\[\]]+`)
	fileName      = flag.String("fileName", "", "File name")
	caseSensitive = flag.Bool("caseSensitive", true, "(optional): True if we consider words case sensitive")
	numOfWords    = flag.Int("numOfWords", 20, "(optional): The number of the words we want to display")
)

func main() {
	flag.Parse()
	if *fileName == "" {
		fmt.Println("File name not specified")
		return
	}

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	numElements := *numOfWords
	useCaseSensitive := *caseSensitive

	wordFreq := hashtable.NewSwissTable[string, int](100)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token := scanner.Text()
		if !useCaseSensitive {
			token = strings.ToLower(token)
		}
		words := wordPattern.FindAllString(token, -1)
		for _, word := range words {
			cnt, _ := wordFreq.Get(word)
			wordFreq.Put(word, cnt+1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	wordCounts := heap.NewMinHeap()
	wordFreq.All(func(word string, count int) bool {
		wordCounts.Push(&heap.Element{Data: word, Priority: count})
		if wordCounts.Len() > numElements {
			wordCounts.Pop()
		}
		return false
	})
	var results []heap.Element
	for ; wordCounts.Len() > 0; wordCounts.Pop() {
		results = append(results, *wordCounts.Top())
	}
	slices.Reverse(results)
	for _, wc := range results {
		fmt.Printf("%d %s\n", wc.Priority, wc.Data)
	}
}

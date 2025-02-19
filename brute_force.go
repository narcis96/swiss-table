package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strings"
// )

// func brute(file *os.File, useCaseSensitive bool, numElements int) {
// 	wordFreq := make(map[string]int)
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		token := scanner.Text()
// 		if !useCaseSensitive {
// 			token = strings.ToLower(token)
// 		}
// 		words := wordPattern.FindAllString(token, -1)
// 		for _, word := range words {
// 			wordFreq[word]++
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}

// 	wordCounts := make([]WordCount, 0, len(wordFreq))
// 	for word, count := range wordFreq {
// 		wordCounts = append(wordCounts, WordCount{word, count})
// 	}

// 	sort.Slice(wordCounts, func(i, j int) bool {
// 		return wordCounts[i].Count > wordCounts[j].Count
// 	})

// 	for i, wc := range wordCounts {
// 		if i >= numElements {
// 			break
// 		}
// 		fmt.Printf("%d %s\n", wc.Count, wc.Word)
// 	}
// }

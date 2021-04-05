package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	t := triangleWords()
	fmt.Printf("\n\nTotal number of triangle words: %d", t)
}

func triangleWords() int {
	tris := triangles(100)
	words := loadWords()
	alphabet := loadAlphabet()

	wordScores := scoreWords(words, alphabet)
	tWords := make([]string, 0)

	for word, score := range wordScores {
		for idx, t := range tris {
			if score == t {
				fmt.Printf("Word %s with score %d is a triangle number at term %d\n", word, score, idx-1)
				tWords = append(tWords, word)
			}
		}
	}

	return len(tWords)
}

func loadAlphabet() map[string]int {
	alphaFile, err := os.Open("./alphabet.csv")
	if err != nil {
		log.Printf("E: %v", err)
	}
	alphaReader := csv.NewReader(alphaFile)
	alphaArr, _ := alphaReader.ReadAll()

	alphabet := make(map[string]int)
	for _, a := range alphaArr {
		i, err := strconv.Atoi(a[1])
		if err != nil {
			log.Printf("E: %v", err)
		}

		alphabet[a[0]] = i
	}
	return alphabet
}

func loadWords() []string {
	file, err := os.Open("./p042_words.txt")
	if err != nil {
		log.Printf("E: %v", err)
	}
	words, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("E: %v", err)
	}
	w := string(words)
	w2 := strings.ReplaceAll(w, "\"", "")
	return strings.Split(w2, ",")
}

func scoreWords(words []string, alphabet map[string]int) map[string]int {
	scores := make(map[string]int)
	for _, word := range words {
		fmt.Printf("Word to score: %s\n", word)
		wordScore := 0
		for _, letter := range word {
			l := string(letter)
			wordScore += alphabet[l]
			fmt.Printf("letter %s has score %d\n", l, alphabet[l])
		}
		scores[word] = wordScore
		fmt.Printf("%s has score: %d\n", word, wordScore)
	}
	return scores
}

func triangles(upper int) (tris []int) {
	for i := 1; i < upper; i++ {
		tris = append(tris, tri(i))
	}
	return
}

func tri(num int) int {
	return num * (num + 1) / 2
}

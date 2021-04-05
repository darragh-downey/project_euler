package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

type NameSorter []string

func (s NameSorter) Len() int           { return len(s) }
func (s NameSorter) Less(i, j int) bool { return s[i] > s[j] }
func (s NameSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	// load the character/value pairs
	alphabet := LoadAlphabet()
	fmt.Printf("%v\n", alphabet)
	// load and sort the names
	namesFile, err := ioutil.ReadFile("./p022_names.txt")
	if err != nil {
		log.Printf("E: %v", err)
	}
	nameBytes := string(namesFile)
	namez := strings.ReplaceAll(nameBytes, "\"", "")
	names := strings.Split(namez, ",")

	sort.Strings(names)

	fmt.Printf("\nFirst element: %s\n", names)
	// score the names
	scores := ScoreNames(names, alphabet)

	// final score
	total := TotalScore(scores)
	fmt.Println(total)
}

func LoadAlphabet() map[string]int {
	alphaFile, err := os.Open("./alphabet.csv")
	if err != nil {
		log.Printf("E: %v", err)
	}
	alphaReader := csv.NewReader(alphaFile)
	alphaArr, _ := alphaReader.ReadAll()
	//fmt.Printf("%s\n", alphaArr)

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

func ScoreNames(names []string, alphabet map[string]int) map[string]int {
	scores := make(map[string]int)

	for idx, name := range names {
		score := 0
		for _, letter := range name {
			score += alphabet[string(letter)]
		}
		scores[name] = score * (idx + 1) // position in 'file'
		fmt.Printf("Idx: %d, Name: %s, Score: %d, Total: %d\n", idx+1, name, score, score*(idx+1))
	}
	return scores
}

func maxName(scores map[string]int) (string, int) {
	max := -1
	maxName := ""
	for name, score := range scores {
		if max < score {
			max = score
			maxName = name
		}
	}
	return maxName, max
}

func TotalScore(scores map[string]int) *big.Int {
	max := big.NewInt(0)
	for _, score := range scores {
		// fmt.Println(score)
		s := big.NewInt(int64(score))
		max.Add(max, s)
	}
	return max
}

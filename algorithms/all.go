package algorithms

import (
	"fmt"
	"log"
	"strings"
	"time"
)

//find largest poly
func isPoly(incomingVal string) bool {
	if incomingVal == "" {
		return false
	}
	lenIncomingVal := len(incomingVal)
	if lenIncomingVal == 1 {
		return false
	}
	midPointIndex := lenIncomingVal / 2
	for i := 0; i < midPointIndex; i++ {
		charToCompare := incomingVal[(lenIncomingVal-1)-i]
		if incomingVal[i] != charToCompare {
			return false
		} else {
			continue
		}
	}
	return true
}

func findLargestPoly(incomingVal string) int {
	if incomingVal == "" {
		return 0
	}
	lenIncomingVal := len(incomingVal)
	if lenIncomingVal == 1 {
		return 0
	}
	var res int
	for i := 0; i < lenIncomingVal-1; i++ {
		for j := lenIncomingVal - 1; j > 0; j-- {
			stringToCheck := incomingVal[i:j]
			stringToCheckLen := len(stringToCheck)
			if isPoly(stringToCheck) {
				if stringToCheckLen > res {
					res = stringToCheckLen
				}
			}
		}
	}
	return res
}

//get count of delta between numbers
func getDeltas(values []int, wantedDelta int) int {
	incomingValLen := len(values)
	if incomingValLen < 2 {
		return 0
	}
	res := 0
	help := make(map[int]struct{}, len(values)*2)
	for i := 0; i < incomingValLen-1; i++ {
		if _, ok := help[values[i]]; ok {
			res++
		}
		help[values[i]+wantedDelta] = struct{}{}
		help[values[i]-wantedDelta] = struct{}{}
	}
	return res
}

func getMin(numA, numB float64) float64 {
	if numA < numB {
		return numA
	}
	return numB
}

//separate array to asc arrays parts
func separateArray() {
	a := []int{5, 32, 8, 34, 1, 3, 45, 56, 5, 2, 33}
	var b [][]int
	stsrtInd := 0
	for i, v := range a {
		if i == 0 {
			continue
		}
		if v <= a[i-1] {
			b = append(b, a[stsrtInd:i])
			stsrtInd = i
		}
		if i == (len(a) - 1) {
			b = append(b, a[stsrtInd:])
		}
	}
	fmt.Println(b)
}

func main() {
	wordStack := []string{"Hi", "Ron", "you", "will", "find", "new", "path"}
	wordChan := make(chan string, 2)
	go reader(wordChan)
	for _, word := range wordStack {
		select {
		case wordChan <- word:
		default:
			log.Print("channel is full")
			wordChan <- word
		}
	}
	close(wordChan)
}

func reader(inChan chan string) {
	for {
		word, ok := <-inChan
		if !ok {
			log.Print("channel is close")
		}
		log.Print(word)
		time.Sleep(1 * time.Second)
	}
}

//////////////////////////////

//fiver
func Regex(pattern, word string) bool {
	hashIndex := strings.Index(pattern, "*")
	if hashIndex == -1 {
		log.Print("1")
		return IsMatchPattern(pattern, word)
	}
	if len(pattern) == 1 {
		log.Print("2")
		return true
	}
	if strings.HasPrefix(pattern, "*") {
		log.Print("3")
		noHashPatternEnd := strings.TrimPrefix(pattern, "*")
		return isEndMatch(noHashPatternEnd, word)
	} else if strings.HasSuffix(pattern, "*") {
		log.Print("4")
		noHashPatternStart := strings.TrimSuffix(pattern, "*")
		return isStartMatch(noHashPatternStart, word)
	}
	log.Print("5")
	patternAsSlice := strings.Split(pattern, "*")
	startMatch := isStartMatch(patternAsSlice[0], word)
	endMatch := isEndMatch(patternAsSlice[1], word)
	return endMatch && startMatch
}

func IsMatchPattern(pattern, word string) bool {
	if len(pattern) != len(word) {
		return false
	}
	for i, char := range pattern {
		if char != int32(word[i]) && char != 63 {
			return false
		}
	}
	return true
}

func isEndMatch(cleanPattern, word string) bool {
	wordIndexToCut := len(word) - len(cleanPattern)
	wordEndToCompare := word[wordIndexToCut:]
	return IsMatchPattern(wordEndToCompare, cleanPattern)
}

func isStartMatch(cleanPattern, word string) bool {
	wordIndexToCut := len(cleanPattern)
	wordStartToCompare := word[:wordIndexToCut]
	return IsMatchPattern(wordStartToCompare, cleanPattern)
}

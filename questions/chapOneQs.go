package questions

import (
	"bytes"
	"log"
)

//1.1
func IsAllUniqueChars(word string) bool {
	wordLen := len(word)
	if wordLen < 2 {
		return true
	}

	for i := 0; i < (wordLen - 1); i++ {
		currChar := word[i]
		restOfChars := word[i+1:]
		log.Printf("here currChar: %v, restOfChars: %+v", currChar, restOfChars)
		for j := 0; j < len(restOfChars); j++ {
			if currChar == restOfChars[j] {
				return false
			}
		}
	}
	return true
}

//1.3
func RemoveDuplicatesChars(word []string) []string {
	wordLen := len(word)
	if wordLen < 2 {
		return word
	}
	//g := []string{"a", "a", "b", "b", "c", "c"}
	tail := 1
	for i := 1; i < wordLen; i++ {
		j := 0
		for j = 0; j <= tail; j++ {
			if word[j] == word[i] {
				break
			}
		}
		// here char is not already appear
		if j == i {
			word[tail] = word[i]
			tail++
		}
	}
	return word[:tail]
}

//1.4
func IsStringsAnagrams(firstString, secondString string) bool {
	firstWordLen := len(firstString)
	secondWordLen := len(secondString)
	if firstWordLen != secondWordLen {
		return false
	}
	firstStringMap := wordToCharsCountMap(firstString)
	secondStringMap := wordToCharsCountMap(secondString)

	for char, firstCount := range firstStringMap {
		if secondCount, existInSecond := secondStringMap[char]; !existInSecond {
			return false
		} else {
			if firstCount != secondCount {
				return false
			}
		}
	}
	return true
}

func wordToCharsCountMap(stringVal string) map[int32]int {
	charsMap := make(map[int32]int, len(stringVal))
	for _, v := range stringVal {
		if vCount, ok := charsMap[v]; !ok {
			vCount = 1
		} else {
			vCount += 1
		}
	}
	return charsMap
}

//1.5
func ReplaceSpaces(stringVal string) string {
	buf := bytes.Buffer{}
	for _, char := range stringVal {
		if char == 32 {
			buf.WriteString("%20")
		} else {
			buf.WriteString(string(char))
		}
	}
	return buf.String()
}

//func try(incoming []string) {
//	incomingLen := len(incoming)
//	tail := 0
//	for i := 0; i < (incomingLen - 1); i++ {
//		current := incoming[i]
//
//		for j := i + 1; j < (incomingLen - 1 - j); j++ {
//			if current == incoming[j] {
//				incoming[j:] = incoming[j+1:]
//			} else {
//
//			}
//		}
//	}
//}

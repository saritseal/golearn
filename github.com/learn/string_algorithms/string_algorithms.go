package string_algorithms

import (
	"strings"
)

func GenerateNGrams(s string, k int, seperator string) []string {
	strArray := []string{}

	splittedString := strings.Split(s, seperator)
	shingleString := []string{}

	for i := 0; i <= (len(splittedString) - k); i++ {

		for _, str := range splittedString[i : i+k] {
			if len(shingleString) == k {
				strArray = append(strArray, strings.Join(shingleString, " "))
				shingleString = nil
			}

			shingleString = append(shingleString, str)
		}

	}

	return strArray
}

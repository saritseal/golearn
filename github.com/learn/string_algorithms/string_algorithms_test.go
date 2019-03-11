package string_algorithms

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestGenerateNGramsFromString(t *testing.T) {

	var output = GenerateNGrams("You can use single workspace but if you want to work with another project out of workspace, you should check your imports. Because when you import golang packages", 1, " ")

	log.Printf("number of tokens identified %d and sample tokens %s", len(output), strings.Join(output, ":"))
}

func TestGenerateNGramsFromFile(t *testing.T) {

	var file, err = os.Open("../../../data/data.txt")
	if err != nil {
		log.Fatal("error while reading the file", err)
	}
	var data, _ = ioutil.ReadAll(file)

	var output = GenerateNGrams(string(data), 1, " ")

	log.Printf("number of tokens identified %d and sample tokens %s", len(output), strings.Join(output[:10], ":"))
}

func BenchmarkGenerateNGramsFromFile(t *testing.B) {

	for n := 0; n < t.N; n++ {
		var file, err = os.Open("../../../data/data.txt")
		if err != nil {
			log.Fatal("error while reading the file", err)
		}
		var data, _ = ioutil.ReadAll(file)

		var output = GenerateNGrams(string(data), 1, " ")

		log.Printf("number of tokens identified %d and sample tokens %s", len(output), strings.Join(output[:10], ":"))
	}

}

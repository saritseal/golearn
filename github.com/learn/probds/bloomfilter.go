package probds

import (
	"log"
	"math/rand"
	"time"
)

type BloomFilterInte interface {
	CreateFilter(int, []HashFunction)
	Add([]byte)
	Check([]byte) bool
}

type HashFunction func([]byte) []byte
type bit int8
type BloomFilterMetaData struct {
	bins          int
	hashFunctions []HashFunction
}

type BloomFilter struct {
	bitArray []bit
}

func CreateFilter(bfMetaData BloomFilterMetaData) *BloomFilter {
	return &BloomFilter{bitArray: make([]bit, 0, bfMetaData.bins)}
}

func evaluateHashFunctions(hashFn []HashFunction) map[int]bit {
	return map[int]bit{1: 1}
}


//Add adds an element to the bloom filter
func (bf *BloomFilter) Add(str []byte) {
	//iterate through the hash functions and identify the bits
	//set the bits from each of the hash functions 
}

// CalcSum this calculates the sum of the characters in the string
func CalcSum(str string) int {
	sum := 0

	for _, x := range str {
		sum = sum + int(x)
	}

	return sum
}

//FindMax finds the maximum element in the array
func FindMax(sums []int) int {

	size := len(sums)

	for c, _ := range sums {

		if c+1 < size {
			if sums[c] > sums[c+1] {
				sums[c], sums[c+1] = sums[c+1], sums[c]
			}
		}
	}
	log.Printf("%#v", sums)

	return sums[size-1]
}

//FindConstants find random constants
func FindConstants(threshold int) (int, int) {
	rand.Seed(time.Now().UnixNano())

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r1.Intn(100), r1.Intn(100)
}

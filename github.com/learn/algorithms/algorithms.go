package algorithms

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type UUID string

func JaccardSimilarity() {
	fmt.Println("jaccard similarity")
}

//GenerateUUID this function generates uuid required for the graph elements
func GenerateUUID() UUID {
	uuidData := make([]byte, 128)
	rand.Read(uuidData)
	uuidData[6] = byte(0x40 | (int(uuidData[6] & 0xf)))
	uuidData[8] = byte(0x80 | (int(uuidData[8] & 0x3f)))
	uuidHex := make([]byte, hex.EncodedLen(len(uuidData)))
	hex.Encode(uuidHex, uuidData)

	return UUID(fmt.Sprintf("%s-%s-%s-%s", uuidHex[0:8], uuidHex[9:13], uuidHex[14:18], uuidHex[19:31]))
}

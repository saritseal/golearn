package algorithms

import (
	"log"
	"testing"	
)

func TestGenerateUUID(t *testing.T) {
	uuid := GenerateUUID()
	log.Printf("UUID Generated %s", uuid)
}

func BenchmarkGenerateUUID(t *testing.B) {
	for i := 0; i < t.N; i++ {
		uuid := GenerateUUID()
		log.Printf("UUID Generated %s", uuid)
	}
}

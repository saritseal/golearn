package probds

import (
	"math/rand"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"gotest.tools/assert"
)

func TestCalcSum(t *testing.T) {
	log.Println(CalcSum("sarit"))
	log.Println(CalcSum("checkdata"))
	log.Println(CalcSum("susmita"))
	log.Println(CalcSum("susmitn"))
}

func TestFindMax(t *testing.T) {
	assert.Equal(t, FindMax([]int{2, 5, 4, 3, 1}), 5)
}

func TestFindConstants(t *testing.T) {
	a, b := FindConstants(5)
	log.Printf("constants %d, %d", a, b)
}

func TestRandom(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	randomValues := make([]int, 0)
	max := 1000000
	for i := 0; i < max; i++ {
		randomValues = append(randomValues, 1+rand.Intn(100))
	}

	log.Printf("RandomValues %v", time.Unix(time.Now().UTC().Unix(), 0))
	log.Printf("data %#v, len %d", randomValues[:max], len(randomValues))
}

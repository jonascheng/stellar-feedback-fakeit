package myfakeit

import (
	"math/rand"

	"github.com/jonascheng/stellar-feedback-fakeit/internal/data"
)

// Check if in lib
func dataCheck(dataVal []string) bool {
	var checkOk bool

	if len(dataVal) == 2 {
		_, checkOk = data.Data[dataVal[0]]
		if checkOk {
			_, checkOk = data.Data[dataVal[0]][dataVal[1]]
		}
	}

	return checkOk
}

// Get Random Value
func getRandValue(r *rand.Rand, dataVal []string) string {
	if !dataCheck(dataVal) {
		return ""
	}
	return data.Data[dataVal[0]][dataVal[1]][r.Intn(len(data.Data[dataVal[0]][dataVal[1]]))]
}

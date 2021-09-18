package util

import (
	"math/rand"
	"time"
)

/*
Random number is seeded to current time due to fact
that random is a sequence and will, by default return same value
*/
func GetRandomNumber(max int, iteration int) int {

	seed := rand.NewSource(time.Now().UnixNano() + int64(iteration))
	randomSeeded := rand.New(seed)

	var randomNumber = randomSeeded.Intn(max)

	return randomNumber
}

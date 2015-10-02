// Copyright © 2015 Christin R. Vozar

package fizzbuzz

import (
	"strconv"
)

// Determine returns Fizz, Buzz, or Fizzbuzz depending on the number provided.
func Determine(i int) string {
	switch {
	case (i % 15) == 0:
		return "FizzBuzz"
	case (i % 5) == 0:
		return "Buzz"
	case (i % 3) == 0:
		return "Fizz"
	default:
		return strconv.Itoa(i)
	}
}

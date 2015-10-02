// Copyright © 2015 Christin R. Vozar

package fizzbuzz

import "testing"

func TestFizz(t *testing.T) {
	v := Determine(9)
	if v != "Fizz" {
		t.Error("Exepected Fizz, got ", v)
	}
}

func TestBuzz(t *testing.T) {
	v := Determine(10)
	if v != "Buzz" {
		t.Error("Exepected Buzz, got ", v)
	}
}

func TestFizzBuzz(t *testing.T) {
	v := Determine(30)
	if v != "FizzBuzz" {
		t.Error("Exepected FizzBuzz, got ", v)
	}
}

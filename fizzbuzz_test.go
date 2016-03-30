// Copyright © 2015 Christin R. Vozar

package fizzbuzz

import "testing"

func TestFizz(t *testing.T) {
	if v := Determine(9); v != "Fizz" {
		t.Error("Exepected Fizz, got ", v)
	}
}

func TestBuzz(t *testing.T) {
	if v := Determine(10); v != "Buzz" {
		t.Error("Exepected Buzz, got ", v)
	}
}

func TestFizzBuzz(t *testing.T) {
	if v := Determine(30); v != "FizzBuzz" {
		t.Error("Exepected FizzBuzz, got ", v)
	}
}

func TestInvalid(t *testing.T) {
	if v := Determine(4); v != "4" {
		t.Error("Exepected 4, got ", v)
	}
}

package main

import (
	"strings"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	if !strings.EqualFold("FizzBuzz", FizzBuzz(15)) {
		t.Errorf("wrong FizzBuzz")
	}
	if !strings.EqualFold("Fizz", FizzBuzz(3)) {
		t.Errorf("wrong Fizz")
	}
	if !strings.EqualFold("Buzz", FizzBuzz(5)) {
		t.Errorf("wrong Buzz")
	}
	if !strings.EqualFold("98", FizzBuzz(98)) {
		t.Errorf("wrong 98")
	}
}

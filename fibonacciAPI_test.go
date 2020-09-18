package main

import (
	"math"
	"testing"
)

// tests fibonacci.next() int overflow error
func TestMaxNumber(t *testing.T) {
	println("Testing int overflow...")
	fTest := fibonacci{math.MaxUint64/2 + 2, math.MaxUint64/2 + 1}
	if n, err := fTest.next(); err == nil {
		t.Error("did not return int overflow error")
	} else if n != math.MaxUint64/2+2 {
		t.Error("current was altered by int overflow, it should remain the same")
	}
	fTest = fibonacci{math.MaxUint64 / 2, math.MaxUint64/2 - 1}
	if _, err := fTest.next(); err != nil {
		t.Errorf("no error was expected but returned: %s", err)
	}
}

// tests fibonacci.previous() returns error when attempting to go below 0
func TestMinNumber(t *testing.T) {
	println("Testing lower bound...")
	fTest := fibonacci{0, 0}
	if n, err := fTest.previous(); err == nil {
		t.Error("did not return error to pass below 0")
	} else if n != 0 {
		t.Errorf("current should 0, but is instead %v", n)
	}
}

// tests whether fibonacci.next() will produce correct numbers of the sequence
func TestFibNext(t *testing.T) {
	println("Testing fibonacci.next() ...")
	fibSeq := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	fTest := fibonacci{0, 0}
	for i := 0; i < 10; i++ {
		if int(fTest.Current) != fibSeq[i] {
			t.Error("next number does not match expected number")
		}
		if _, err := fTest.next(); err != nil {
			t.Errorf("no error was expected but returned: %s", err)
		}
	}
}

// tests whether fibonacci.previous() will produce correct numbers of the sequence
func TestFibPrevious(t *testing.T) {
	println("Testing fibonacci.previous() ...")
	fibSeq := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	fTest := fibonacci{55, 34}
	for i := 0; i < 10; i++ {
		if int(fTest.Current) != fibSeq[len(fibSeq)-1-i] {
			t.Error("previous number does not match expected number")
		}
		if _, err := fTest.previous(); err != nil {
			t.Errorf("no error was expected but returned: %s", err)
		}
	}
}

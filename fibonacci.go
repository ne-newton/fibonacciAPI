package main

import (
	"errors"
	"math"
)

type fibonacci struct {
	Current  int
	Previous int
}

// returns the current number in the fibonacci sequence, with no alterations
func (f *fibonacci) current() int {
	return f.Current
}

// replaces the current number in the fibonacci sequence with the current + previous
// replaces the previous number with the former current number, then returns the new current
// if both current and previous are 0, current is changed to 1 to start sequence
func (f *fibonacci) next() (int, error) {
	if f.Current > (math.MaxInt32 - f.Previous) {
		return f.Current, errors.New("fibonacci: int overflow")
	}
	if !(f.Current == 0 && f.Previous == 0) {
		f.Current, f.Previous = f.Current+f.Previous, f.Current
		return f.Current, nil
	}
	f.Current = 1
	return f.Current, nil
}

// replaces the current number with the previous number from the fibonacci sequence
// replaces the previous number with former current - previous, returns new current number
// if current is already 0, there is no previous number, so error is returned along with current as 0
func (f *fibonacci) previous() (int, error) {
	if f.Current != 0 {
		f.Current, f.Previous = f.Previous, f.Current-f.Previous
		return f.Current, nil
	}
	return 0, errors.New("fibonacci: no previous past 0")
}

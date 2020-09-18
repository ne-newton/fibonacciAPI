package main

import (
	"errors"
	"math"
)

type fibonacci struct {
	Current  uint64
	Previous uint64
}

// returns the current number in the fibonacci sequence, with no alterations
func (f *fibonacci) current() uint64 {
	return f.Current
}

// replaces the current number in the fibonacci sequence with the current + previous
// replaces the previous number with the former current number, then returns the new current
// if both current and previous are 0, current is changed to 1 to start sequence
// will return overflow error if next would exceed uint64 size
func (f *fibonacci) next() (uint64, error) {
	if f.Current > (math.MaxUint64 - f.Previous) {
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
func (f *fibonacci) previous() (uint64, error) {
	if f.Current != 0 {
		f.Current, f.Previous = f.Previous, f.Current-f.Previous
		return f.Current, nil
	}
	return 0, errors.New("fibonacci: no previous past 0")
}

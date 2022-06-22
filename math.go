// Copyright (c) 2022 Matt Way
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE THE SOFTWARE.

// Package math provides math-related utilities.
package math

import (
	"golang.org/x/exp/constraints"
)

// Numeric is a constraint that permits any integer or floating point number.
type Numeric interface {
	constraints.Integer | constraints.Float
}

// SignedNumeric is a constraint that permits any signed integer or floating
// point number.
type SignedNumeric interface {
	constraints.Signed | constraints.Float
}

// Abs returns the absolute value of the given signed number.
func Abs[T SignedNumeric](x T) T {
	if x < 0 {
		return -x
	}

	return x
}

// Min returns the minimum value of the two given numbers.
func Min[T Numeric](x T, y T) T {
	if x < y {
		return x
	}

	return y
}

// MinN returns the minimum value of the given numbers.
func MinN[T Numeric](x ...T) T {
	var min T

	if len(x) == 0 {
		return min
	}

	min = x[0]
	for _, n := range x[1:] {
		if n < min {
			min = n
		}
	}

	return min
}

// Max returns the maximum value of the two given numbers.
func Max[T Numeric](x T, y T) T {
	if x > y {
		return x
	}

	return y
}

// MaxN returns the maximum value of the given numbers.
func MaxN[T Numeric](x ...T) T {
	var max T

	if len(x) == 0 {
		return max
	}

	max = x[0]
	for _, n := range x[1:] {
		if n > max {
			max = n
		}
	}

	return max
}

// Mean returns the truncated average value of all given numbers.
func Mean[T Numeric](x ...T) T {
	var (
		size  = len(x)
		total T
	)

	for len(x) >= 8 {
		total += x[0] + x[1] + x[2] + x[3] + x[4] + x[5] + x[6] + x[7]
		x = x[8:]
	}

	for _, n := range x {
		total += n
	}

	return total / T(size)
}

// MeanFloat64 returns the average value of all given numbers.
func MeanFloat64[T Numeric](x ...T) float64 {
	var (
		size  = len(x)
		total T
	)

	for len(x) >= 8 {
		total += x[0] + x[1] + x[2] + x[3] + x[4] + x[5] + x[6] + x[7]
		x = x[8:]
	}

	for _, n := range x {
		total += n
	}

	return float64(total) / float64(size)
}

// Clamp clamps the given value between [min,max] (inclusive).
func Clamp[T Numeric](x T, min T, max T) T {
	switch {
	case x > max:
		return max
	case x < min:
		return min
	default:
		return x
	}
}

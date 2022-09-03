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
	"math"
	_ "unsafe" // for go:linkname

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

// Unsigned32 is a constraint that permits any type that can fit a 32-bit
// unsigned integer.
type Unsigned32 interface {
	~int | ~int64 | ~uint | ~uint32 | ~uint64 | constraints.Float
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
	return MinN(x, y)
}

// MinN returns the minimum value of the given numbers.
func MinN[T Numeric](x ...T) T {
	switch len(x) {
	case 0:
		return 0
	case 1:
		return x[0]
	case 2:
		if x[0] < x[1] {
			return x[0]
		}
		return x[1]
	case 3:
		y := x[0]
		if x[1] < y {
			y = x[1]
		}
		if x[2] < y {
			y = x[2]
		}
		return y
	case 4:
		y := x[0]
		if x[1] < y {
			y = x[1]
		}
		if x[2] < y {
			y = x[2]
		}
		if x[3] < y {
			y = x[3]
		}
		return y
	default:
		min := x[0]
		for _, n := range x[1:] {
			if n < min {
				min = n
			}
		}

		return min
	}
}

// Max returns the maximum value of the two given numbers.
func Max[T Numeric](x T, y T) T {
	return MaxN(x, y)
}

// MaxN returns the maximum value of the given numbers.
func MaxN[T Numeric](x ...T) T {
	switch len(x) {
	case 0:
		return 0
	case 1:
		return x[0]
	case 2:
		if x[0] > x[1] {
			return x[0]
		}
		return x[1]
	case 3:
		y := x[0]
		if x[1] > y {
			y = x[1]
		}
		if x[2] > y {
			y = x[2]
		}
		return y
	case 4:
		y := x[0]
		if x[1] > y {
			y = x[1]
		}
		if x[2] > y {
			y = x[2]
		}
		if x[3] > y {
			y = x[3]
		}
		return y
	default:
		max := x[0]
		for _, n := range x[1:] {
			if n > max {
				max = n
			}
		}

		return max
	}
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

// ClampMin clamps the given value such that it is not less than min (inclusive).
func ClampMin[T Numeric](x T, min T) T {
	switch {
	case x < min:
		return min
	default:
		return x
	}
}

// ClampMax clamps the given value such that it is not greater than max
// (inclusive).
func ClampMax[T Numeric](x T, max T) T {
	switch {
	case x > max:
		return max
	default:
		return x
	}
}

// NextPowerOf2 returns the next T greater than x that is a power of 2. If x is
// a power of 2 itself, x is returned.
func NextPowerOf2[T Numeric](x T) T {
	if x <= 1 {
		return 1
	}

	var (
		l2 = math.Log2(float64(x))
		hi = T(int(1) << int(math.Ceil(l2)))
	)

	return hi
}

// ClosestPowerOf2 returns the next T greater than x that is a power of 2.
func ClosestPowerOf2[T Numeric](x T) T {
	if x <= 1 {
		return 1
	}

	var (
		l2 = math.Log2(float64(x))
		lo = T(int(1) << int(math.Floor(l2)))
		hi = T(int(1) << int(math.Ceil(l2)))
	)

	if hi-x > x-lo {
		return lo
	}

	return hi
}

// Precision truncates x to the given precision. If precision is < 0, x is
// returned unchanged; if x == 0, x is rounded to the nearest integer;
// otherwise, the precision of x is changed and the resulting value rounded.
func Precision[T constraints.Float](x T, precision int) T {
	if precision < 0 {
		return x
	}

	if precision == 0 {
		return T(math.Round(float64(x)))
	}

	var (
		coeff = math.Pow(10.0, float64(precision))
		tmp   = math.Round(float64(x) * coeff)
	)

	return T(tmp / coeff)
}

// Fastrand returns a pseudorandom T in the range [0, 1<<32-1).
func Fastrand[T Unsigned32]() T {
	return T(fastrand())
}

// Fastrandn returns a pseudorandom T in the range [0, min(x, 1<<32-1)). If x
// is greater than 1<<32-1, it will be clamped.
func Fastrandn[T Unsigned32](x T) T {
	return T(fastrandn(uint32(Clamp(x, 0, math.MaxUint32))))
}

//go:linkname fastrand runtime.fastrand
func fastrand() uint32

//go:linkname fastrandn runtime.fastrandn
func fastrandn(uint32) uint32

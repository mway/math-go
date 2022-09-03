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

package math_test

import (
	"strconv"
	"testing"

	"go.mway.dev/math"
)

func BenchmarkAbs(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		math.Abs(-i)
	}
}

func BenchmarkMin(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		math.Min(i+1, i)
	}
}

func BenchmarkMinN(b *testing.B) {
	b.Run("2 args", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			math.MinN(i+1, i)
		}
	})

	b.Run("3 args", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			math.MinN(i+2, i+1, i)
		}
	})

	b.Run("4 args", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			math.MinN(i+3, i+2, i+1, i)
		}
	})

	b.Run("5 args", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			math.MinN(i+4, i+3, i+2, i+1, i)
		}
	})
}

func BenchmarkMax(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		math.Max(i, i+1)
	}
}

func BenchmarkMaxN(b *testing.B) {
	b.Run("2 args", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			math.MaxN(i, i+1)
		}
	})

	b.Run("3 args", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			math.MaxN(i, i+1, i+2)
		}
	})

	b.Run("4 args", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			math.MaxN(i, i+1, i+2, i+3)
		}
	})

	b.Run("5 args", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			math.MaxN(i, i+1, i+2, i+3, i+4)
		}
	})
}

func BenchmarkMean(b *testing.B) {
	sizes := []int{
		2 << 0,
		2<<1 - 1,
		2<<2 - 1,
		2<<3 - 1,
		2<<4 - 1,
		2<<5 - 1,
		2<<6 - 1,
	}

	for _, size := range sizes {
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			numbers := make([]uint64, size)
			for i := 0; i < len(numbers); i++ {
				numbers[i] = uint64(i)
			}

			b.Run("typed", func(b *testing.B) {
				b.ReportAllocs()

				for i := 0; i < b.N; i++ {
					math.Mean(numbers...)
				}
			})

			b.Run("float64", func(b *testing.B) {
				b.ReportAllocs()

				for i := 0; i < b.N; i++ {
					math.MeanFloat64(numbers...)
				}
			})
		})
	}
}

func BenchmarkNextPowerOf2(b *testing.B) {
	pow2s := make([]int, 8)
	for i := range pow2s {
		pow2s[i] = 1 << (i * 2)
	}

	for _, give := range pow2s {
		b.Run(strconv.Itoa(give), func(b *testing.B) {
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				math.NextPowerOf2(give)
			}
		})
	}
}

func BenchmarkClosestPowerOf2(b *testing.B) {
	pow2s := make([]int, 8)
	for i := range pow2s {
		pow2s[i] = 1 << (i * 2)
	}

	for _, give := range pow2s {
		b.Run(strconv.Itoa(give), func(b *testing.B) {
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				math.ClosestPowerOf2(give)
			}
		})
	}
}

func BenchmarkFastrand(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		math.Fastrand[int]()
	}
}

func BenchmarkFastrandn(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		math.Fastrandn(1024)
	}
}

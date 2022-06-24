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
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mway.dev/math"
)

func TestAbs(t *testing.T) {
	require.Equal(t, 10, math.Abs(10))
	require.Equal(t, int(10), math.Abs(-10))
	require.Equal(t, int8(10), math.Abs[int8](10))
	require.Equal(t, int8(10), math.Abs[int8](-10))
	require.Equal(t, int16(10), math.Abs[int16](10))
	require.Equal(t, int16(10), math.Abs[int16](-10))
	require.Equal(t, int32(10), math.Abs[int32](10))
	require.Equal(t, int32(10), math.Abs[int32](-10))
	require.Equal(t, int64(10), math.Abs[int64](10))
	require.Equal(t, int64(10), math.Abs[int64](-10))
	require.Equal(t, time.Duration(10), math.Abs[time.Duration](-10))
	require.Equal(t, float32(10.0), math.Abs[float32](10.0))
	require.Equal(t, float32(10.0), math.Abs[float32](-10.0))
	require.Equal(t, float64(10.0), math.Abs(-10.0))
	require.Equal(t, float64(10.0), math.Abs(10.0))
}

func TestMin(t *testing.T) {
	require.Equal(t, 10, math.Min(100, 10))
	require.Equal(t, 10, math.Min(10, 100))
	require.Equal(t, int8(10), math.Min[int8](100, 10))
	require.Equal(t, int8(10), math.Min[int8](10, 100))
	require.Equal(t, int16(10), math.Min[int16](100, 10))
	require.Equal(t, int16(10), math.Min[int16](10, 100))
	require.Equal(t, int32(10), math.Min[int32](100, 10))
	require.Equal(t, int32(10), math.Min[int32](10, 100))
	require.Equal(t, int64(10), math.Min[int64](100, 10))
	require.Equal(t, int64(10), math.Min[int64](10, 100))
	require.Equal(t, uint(10), math.Min[uint](100, 10))
	require.Equal(t, uint(10), math.Min[uint](10, 100))
	require.Equal(t, uint8(10), math.Min[uint8](100, 10))
	require.Equal(t, uint8(10), math.Min[uint8](10, 10))
	require.Equal(t, uint16(10), math.Min[uint16](100, 10))
	require.Equal(t, uint16(10), math.Min[uint16](10, 100))
	require.Equal(t, uint32(10), math.Min[uint32](100, 10))
	require.Equal(t, uint32(10), math.Min[uint32](10, 100))
	require.Equal(t, uint64(10), math.Min[uint64](100, 10))
	require.Equal(t, uint64(10), math.Min[uint64](10, 100))
	require.Equal(t, time.Duration(10), math.Min[time.Duration](10, 100))
	require.Equal(t, float32(10.0), math.Min[float32](10.0, 100.0))
	require.Equal(t, float32(10.0), math.Min[float32](100.0, 10.0))
	require.Equal(t, float64(10.0), math.Min(10.0, 100.0))
	require.Equal(t, float64(10.0), math.Min(100.0, 10.0))
}

func TestMinN(t *testing.T) {
	require.Equal(t, 0, math.MinN[int]())
	require.Equal(t, 10, math.MinN(10))
	require.Equal(t, 10, math.MinN(100, 50, 10))
	require.Equal(t, 10, math.MinN(10, 50, 100))
	require.Equal(t, int8(10), math.MinN[int8](100, 50, 10))
	require.Equal(t, int8(10), math.MinN[int8](10, 50, 100))
	require.Equal(t, int16(10), math.MinN[int16](100, 50, 10))
	require.Equal(t, int16(10), math.MinN[int16](10, 50, 100))
	require.Equal(t, int32(10), math.MinN[int32](100, 50, 10))
	require.Equal(t, int32(10), math.MinN[int32](10, 50, 100))
	require.Equal(t, int64(10), math.MinN[int64](100, 50, 10))
	require.Equal(t, int64(10), math.MinN[int64](10, 50, 100))
	require.Equal(t, uint(10), math.MinN[uint](100, 50, 10))
	require.Equal(t, uint(10), math.MinN[uint](10, 50, 100))
	require.Equal(t, uint8(10), math.MinN[uint8](100, 50, 10))
	require.Equal(t, uint8(10), math.MinN[uint8](10, 10))
	require.Equal(t, uint16(10), math.MinN[uint16](100, 50, 10))
	require.Equal(t, uint16(10), math.MinN[uint16](10, 50, 100))
	require.Equal(t, uint32(10), math.MinN[uint32](100, 50, 10))
	require.Equal(t, uint32(10), math.MinN[uint32](10, 50, 100))
	require.Equal(t, uint64(10), math.MinN[uint64](100, 50, 10))
	require.Equal(t, uint64(10), math.MinN[uint64](10, 50, 100))
	require.Equal(t, time.Duration(10), math.MinN[time.Duration](10, 50, 100))
	require.Equal(t, float32(10.0), math.MinN[float32](10.0, 50.0, 100.0))
	require.Equal(t, float32(10.0), math.MinN[float32](100.0, 50.0, 10.0))
	require.Equal(t, float64(10.0), math.MinN(10.0, 50.0, 100.0))
	require.Equal(t, float64(10.0), math.MinN(100.0, 50.0, 10.0))
}

func TestMax(t *testing.T) {
	require.Equal(t, 0, math.MaxN[int]())
	require.Equal(t, 10, math.MaxN(10))
	require.Equal(t, 100, math.Max(10, 100))
	require.Equal(t, 100, math.Max(100, 100))
	require.Equal(t, int8(100), math.Max[int8](10, 100))
	require.Equal(t, int8(100), math.Max[int8](100, 10))
	require.Equal(t, int16(100), math.Max[int16](10, 100))
	require.Equal(t, int16(100), math.Max[int16](100, 10))
	require.Equal(t, int32(100), math.Max[int32](10, 100))
	require.Equal(t, int32(100), math.Max[int32](100, 10))
	require.Equal(t, int64(100), math.Max[int64](10, 100))
	require.Equal(t, int64(100), math.Max[int64](100, 10))
	require.Equal(t, uint(100), math.Max[uint](10, 100))
	require.Equal(t, uint(100), math.Max[uint](100, 10))
	require.Equal(t, uint8(100), math.Max[uint8](10, 100))
	require.Equal(t, uint8(100), math.Max[uint8](100, 10))
	require.Equal(t, uint16(100), math.Max[uint16](10, 100))
	require.Equal(t, uint16(100), math.Max[uint16](100, 10))
	require.Equal(t, uint32(100), math.Max[uint32](10, 100))
	require.Equal(t, uint32(100), math.Max[uint32](100, 10))
	require.Equal(t, uint64(100), math.Max[uint64](10, 100))
	require.Equal(t, uint64(100), math.Max[uint64](100, 10))
	require.Equal(t, time.Duration(100), math.Max[time.Duration](100, 10))
	require.Equal(t, float32(100.0), math.Max[float32](10.0, 100.0))
	require.Equal(t, float32(100.0), math.Max[float32](100.0, 10.0))
	require.Equal(t, float64(100.0), math.Max(10.0, 100.0))
	require.Equal(t, float64(100.0), math.Max(100.0, 10.0))
}

func TestMaxN(t *testing.T) {
	require.Equal(t, 100, math.MaxN(10, 50, 100))
	require.Equal(t, 100, math.MaxN(100, 50, 100))
	require.Equal(t, int8(100), math.MaxN[int8](10, 50, 100))
	require.Equal(t, int8(100), math.MaxN[int8](100, 50, 10))
	require.Equal(t, int16(100), math.MaxN[int16](10, 50, 100))
	require.Equal(t, int16(100), math.MaxN[int16](100, 50, 10))
	require.Equal(t, int32(100), math.MaxN[int32](10, 50, 100))
	require.Equal(t, int32(100), math.MaxN[int32](100, 50, 10))
	require.Equal(t, int64(100), math.MaxN[int64](10, 50, 100))
	require.Equal(t, int64(100), math.MaxN[int64](100, 50, 10))
	require.Equal(t, uint(100), math.MaxN[uint](10, 50, 100))
	require.Equal(t, uint(100), math.MaxN[uint](100, 50, 10))
	require.Equal(t, uint8(100), math.MaxN[uint8](10, 50, 100))
	require.Equal(t, uint8(100), math.MaxN[uint8](100, 50, 10))
	require.Equal(t, uint16(100), math.MaxN[uint16](10, 50, 100))
	require.Equal(t, uint16(100), math.MaxN[uint16](100, 50, 10))
	require.Equal(t, uint32(100), math.MaxN[uint32](10, 50, 100))
	require.Equal(t, uint32(100), math.MaxN[uint32](100, 50, 10))
	require.Equal(t, uint64(100), math.MaxN[uint64](10, 50, 100))
	require.Equal(t, uint64(100), math.MaxN[uint64](100, 50, 10))
	require.Equal(t, time.Duration(100), math.MaxN[time.Duration](100, 50, 10))
	require.Equal(t, float32(100.0), math.MaxN[float32](10.0, 50.0, 100.0))
	require.Equal(t, float32(100.0), math.MaxN[float32](100.0, 50.0, 10.0))
	require.Equal(t, float64(100.0), math.MaxN(10.0, 50.0, 100.0))
	require.Equal(t, float64(100.0), math.MaxN(100.0, 50.0, 10.0))
}

func TestMean(t *testing.T) {
	require.Equal(t, 3, math.Mean(1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, int8(3), math.Mean[int8](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, int16(3), math.Mean[int16](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, int32(3), math.Mean[int32](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, int64(3), math.Mean[int64](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, uint(3), math.Mean[uint](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, uint8(3), math.Mean[uint8](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, uint16(3), math.Mean[uint16](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, uint32(3), math.Mean[uint32](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, uint64(3), math.Mean[uint64](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, time.Duration(3), math.Mean[time.Duration](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, float32(3.5), math.Mean[float32](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, float32(3.5), math.Mean[float32](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, float64(3.5), math.Mean[float64](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, float64(3.5), math.Mean[float64](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
}

func TestMeanFloat64(t *testing.T) {
	require.Equal(t, 3.5, math.MeanFloat64(1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, 3.5, math.MeanFloat64[int8](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, 3.5, math.MeanFloat64[int16](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, 3.5, math.MeanFloat64[int32](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, 3.5, math.MeanFloat64[int64](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, 3.5, math.MeanFloat64[uint](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, 3.5, math.MeanFloat64[uint8](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, 3.5, math.MeanFloat64[uint16](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, 3.5, math.MeanFloat64[uint32](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, 3.5, math.MeanFloat64[uint64](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, 3.5, math.MeanFloat64[time.Duration](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, float64(3.5), math.MeanFloat64[float32](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, float64(3.5), math.MeanFloat64[float32](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, float64(3.5), math.MeanFloat64[float64](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
	require.Equal(t, float64(3.5), math.MeanFloat64[float64](1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6))
}

func TestClamp(t *testing.T) {
	require.Equal(t, 5, math.Clamp(5, 1, 10))
	require.Equal(t, 10, math.Clamp(15, 1, 10))
	require.Equal(t, 1, math.Clamp(-1, 1, 10))
	require.Equal(t, 1.5, math.Clamp(1.5, 1.0, 2.0))
	require.Equal(t, time.Second, math.Clamp(time.Hour, time.Millisecond, time.Second))
}

func TestNextPowerOf2(t *testing.T) {
	cases := [][2]int{
		// give, want
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 4},
		{5, 8},
		{6, 8},
		{7, 8},
		{8, 8},
		{9, 16},
		{10, 16},
		{11, 16},
		{12, 16},
		{13, 16},
		{14, 16},
		{15, 16},
		{16, 16},
		{17, 32},
		{18, 32},
		{19, 32},
		{20, 32},
		{21, 32},
		{22, 32},
		{23, 32},
		{24, 32},
		{25, 32},
		{26, 32},
		{27, 32},
		{28, 32},
		{29, 32},
		{30, 32},
		{31, 32},
		{32, 32},
	}

	for _, pair := range cases {
		require.Equal(t, pair[1], math.NextPowerOf2(pair[0]))
		require.Equal(t, int8(pair[1]), math.NextPowerOf2(int8(pair[0])))
		require.Equal(t, int16(pair[1]), math.NextPowerOf2(int16(pair[0])))
		require.Equal(t, int32(pair[1]), math.NextPowerOf2(int32(pair[0])))
		require.Equal(t, int64(pair[1]), math.NextPowerOf2(int64(pair[0])))
		require.Equal(t, uint(pair[1]), math.NextPowerOf2(uint(pair[0])))
		require.Equal(t, uint8(pair[1]), math.NextPowerOf2(uint8(pair[0])))
		require.Equal(t, uint16(pair[1]), math.NextPowerOf2(uint16(pair[0])))
		require.Equal(t, uint32(pair[1]), math.NextPowerOf2(uint32(pair[0])))
		require.Equal(t, uint64(pair[1]), math.NextPowerOf2(uint64(pair[0])))
		require.Equal(t, float32(pair[1]), math.NextPowerOf2(float32(pair[0])))
		require.Equal(t, float64(pair[1]), math.NextPowerOf2(float64(pair[0])))
		require.Equal(t, time.Duration(pair[1]), math.NextPowerOf2(time.Duration(pair[0])))
	}
}

func TestClosestPowerOf2(t *testing.T) {
	cases := [][2]int{
		// give, want
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 4},
		{5, 4},
		{6, 8},
		{7, 8},
		{8, 8},
		{9, 8},
		{10, 8},
		{11, 8},
		{12, 16},
		{13, 16},
		{14, 16},
		{15, 16},
		{16, 16},
		{17, 16},
		{18, 16},
		{19, 16},
		{20, 16},
		{21, 16},
		{22, 16},
		{23, 16},
		{24, 32},
		{25, 32},
		{26, 32},
		{27, 32},
		{28, 32},
		{29, 32},
		{30, 32},
		{31, 32},
		{32, 32},
	}

	for _, pair := range cases {
		require.Equal(t, pair[1], math.ClosestPowerOf2(pair[0]))
		require.Equal(t, int8(pair[1]), math.ClosestPowerOf2(int8(pair[0])))
		require.Equal(t, int16(pair[1]), math.ClosestPowerOf2(int16(pair[0])))
		require.Equal(t, int32(pair[1]), math.ClosestPowerOf2(int32(pair[0])))
		require.Equal(t, int64(pair[1]), math.ClosestPowerOf2(int64(pair[0])))
		require.Equal(t, uint(pair[1]), math.ClosestPowerOf2(uint(pair[0])))
		require.Equal(t, uint8(pair[1]), math.ClosestPowerOf2(uint8(pair[0])))
		require.Equal(t, uint16(pair[1]), math.ClosestPowerOf2(uint16(pair[0])))
		require.Equal(t, uint32(pair[1]), math.ClosestPowerOf2(uint32(pair[0])))
		require.Equal(t, uint64(pair[1]), math.ClosestPowerOf2(uint64(pair[0])))
		require.Equal(t, float32(pair[1]), math.ClosestPowerOf2(float32(pair[0])))
		require.Equal(t, float64(pair[1]), math.ClosestPowerOf2(float64(pair[0])))
		require.Equal(t, time.Duration(pair[1]), math.ClosestPowerOf2(time.Duration(pair[0])))
	}
}

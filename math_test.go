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

	"github.com/stretchr/testify/require"
	"go.mway.dev/math"
)

func TestMin(t *testing.T) {
	require.Equal(t, int(10), math.Min[int](100, 10))
	require.Equal(t, int8(10), math.Min[int8](100, 10))
	require.Equal(t, int16(10), math.Min[int16](100, 10))
	require.Equal(t, int32(10), math.Min[int32](100, 10))
	require.Equal(t, int64(10), math.Min[int64](100, 10))
	require.Equal(t, uint(10), math.Min[uint](100, 10))
	require.Equal(t, uint8(10), math.Min[uint8](100, 10))
	require.Equal(t, uint16(10), math.Min[uint16](100, 10))
	require.Equal(t, uint32(10), math.Min[uint32](100, 10))
	require.Equal(t, uint64(10), math.Min[uint64](100, 10))
}

func TestMax(t *testing.T) {
	require.Equal(t, int(100), math.Max[int](10, 100))
	require.Equal(t, int8(100), math.Max[int8](10, 100))
	require.Equal(t, int16(100), math.Max[int16](10, 100))
	require.Equal(t, int32(100), math.Max[int32](10, 100))
	require.Equal(t, int64(100), math.Max[int64](10, 100))
	require.Equal(t, uint(100), math.Max[uint](10, 100))
	require.Equal(t, uint8(100), math.Max[uint8](10, 100))
	require.Equal(t, uint16(100), math.Max[uint16](10, 100))
	require.Equal(t, uint32(100), math.Max[uint32](10, 100))
	require.Equal(t, uint64(100), math.Max[uint64](10, 100))
}

func TestMean(t *testing.T) {
	require.Equal(t, int(2), math.Mean[int](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, int8(2), math.Mean[int8](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, int16(2), math.Mean[int16](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, int32(2), math.Mean[int32](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, int64(2), math.Mean[int64](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, uint(2), math.Mean[uint](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, uint8(2), math.Mean[uint8](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, uint16(2), math.Mean[uint16](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, uint32(2), math.Mean[uint32](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, uint64(2), math.Mean[uint64](1, 1, 2, 2, 3, 3, 4, 4))
}

func TestMeanFloat64(t *testing.T) {
	require.Equal(t, 2.5, math.MeanFloat64[int](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, 2.5, math.MeanFloat64[int8](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, 2.5, math.MeanFloat64[int16](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, 2.5, math.MeanFloat64[int32](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, 2.5, math.MeanFloat64[int64](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, 2.5, math.MeanFloat64[uint](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, 2.5, math.MeanFloat64[uint8](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, 2.5, math.MeanFloat64[uint16](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, 2.5, math.MeanFloat64[uint32](1, 1, 2, 2, 3, 3, 4, 4))
	require.Equal(t, 2.5, math.MeanFloat64[uint64](1, 1, 2, 2, 3, 3, 4, 4))
}

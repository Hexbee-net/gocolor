// Copyright © 2019 Xavier Basty <xavier@hexbee.net>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gocolor_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Hexbee-net/gocolor"
)

const precision = 1e-8

type ConversionTest struct {
	from []float64
	to   []float64
}

////////////////////////////////////////

func TestFromRGB(t *testing.T) {
	t.Run("convert RGB color to HSL", TestRGBtoHSL)
	t.Run("convert RGB color to HSV", TestRGBtoHSV)
	t.Run("convert RGB color to YIQ", TestRGBtoYIQ)
	t.Run("convert RGB color to YUV", TestRGBtoYUV)
	t.Run("convert RGB color to CMY", TestRGBtoCMY)
	t.Run("convert RGB color to HTML", TestRGBtoHEX)
	t.Run("convert RGB color to XYZ", TestRGBtoXYZ)
}

func TestRGBtoHSL_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-1, 0, 0},
		{2, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.RGBtoHSL(tests[n][0], tests[n][1], tests[n][2])
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestRGBtoHSL(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{0.00, 0.00, 0.00}, to: []float64{0.000, 0.00, 0.000}},
		{from: []float64{1.00, 1.00, 1.00}, to: []float64{0.000, 0.00, 1.000}},
		{from: []float64{1.00, 0.00, 0.00}, to: []float64{0.000, 1.00, 0.500}},
		{from: []float64{0.00, 1.00, 0.00}, to: []float64{120.0, 1.00, 0.500}},
		{from: []float64{0.00, 0.00, 1.00}, to: []float64{240.0, 1.00, 0.500}},
		{from: []float64{1.00, 1.00, 0.00}, to: []float64{60.00, 1.00, 0.500}},
		{from: []float64{1.00, 0.00, 1.00}, to: []float64{300.0, 1.00, 0.500}},
		{from: []float64{0.00, 1.00, 1.00}, to: []float64{180.0, 1.00, 0.500}},
		{from: []float64{1.00, 0.50, 0.00}, to: []float64{30.00, 1.00, 0.500}},
		{from: []float64{1.00, 0.50, 0.50}, to: []float64{0.000, 1.00, 0.750}},
		{from: []float64{0.50, 1.00, 0.00}, to: []float64{90.00, 1.00, 0.500}},
		{from: []float64{0.50, 1.00, 0.50}, to: []float64{120.0, 1.00, 0.750}},
		{from: []float64{0.50, 0.00, 1.00}, to: []float64{270.0, 1.00, 0.500}},
		{from: []float64{0.50, 0.50, 1.00}, to: []float64{240.0, 1.00, 0.750}},
		{from: []float64{0.50, 0.50, 0.50}, to: []float64{0.000, 0.00, 0.500}},
		{from: []float64{0.50, 0.50, 0.00}, to: []float64{60.00, 1.00, 0.250}},
		{from: []float64{0.50, 0.00, 0.50}, to: []float64{300.0, 1.00, 0.250}},
		{from: []float64{0.00, 0.50, 0.50}, to: []float64{180.0, 1.00, 0.250}},
		{from: []float64{0.50, 0.00, 0.00}, to: []float64{0.000, 1.00, 0.250}},
		{from: []float64{0.00, 0.50, 0.00}, to: []float64{120.0, 1.00, 0.250}},
		{from: []float64{0.00, 0.00, 0.50}, to: []float64{240.0, 1.00, 0.250}},
		{from: []float64{1.00, 0.50, 0.25}, to: []float64{20.00, 1.00, 0.625}},
		{from: []float64{0.50, 1.00, 0.25}, to: []float64{100.0, 1.00, 0.625}},
		{from: []float64{0.50, 0.25, 1.00}, to: []float64{260.0, 1.00, 0.625}},
		{from: []float64{0.25, 1.00, 0.50}, to: []float64{140.0, 1.00, 0.625}},
		{from: []float64{0.25, 0.50, 1.00}, to: []float64{220.0, 1.00, 0.625}},
		{from: []float64{0.75, 0.00, 0.00}, to: []float64{0.000, 1.00, 0.375}},
		{from: []float64{0.00, 0.75, 0.00}, to: []float64{120.0, 1.00, 0.375}},
		{from: []float64{0.00, 0.00, 0.75}, to: []float64{240.0, 1.00, 0.375}},
		{from: []float64{0.75, 0.75, 0.00}, to: []float64{60.00, 1.00, 0.375}},
		{from: []float64{0.75, 0.00, 0.75}, to: []float64{300.0, 1.00, 0.375}},
		{from: []float64{0.00, 0.75, 0.75}, to: []float64{180.0, 1.00, 0.375}},
		{from: []float64{0.75, 0.50, 0.00}, to: []float64{40.00, 1.00, 0.375}},
		{from: []float64{0.75, 0.00, 0.50}, to: []float64{320.0, 1.00, 0.375}},
		{from: []float64{0.75, 0.50, 0.50}, to: []float64{0.000, 1.0 / 3.0, 0.625}},
		{from: []float64{0.50, 0.75, 0.00}, to: []float64{80.00, 1.00, 0.375}},
		{from: []float64{0.00, 0.75, 0.50}, to: []float64{160.0, 1.00, 0.375}},
		{from: []float64{0.50, 0.75, 0.50}, to: []float64{120.0, 1.0 / 3.0, 0.625}},
		{from: []float64{0.50, 0.00, 0.75}, to: []float64{280.0, 1.00, 0.375}},
		{from: []float64{0.00, 0.50, 0.75}, to: []float64{200.0, 1.00, 0.375}},
		{from: []float64{0.50, 0.50, 0.75}, to: []float64{240.0, 1.0 / 3.0, 0.625}},
	}

	for n := 0; n < len(tests); n++ {
		h, s, l, err := gocolor.RGBtoHSL(tests[n].from[0], tests[n].from[1], tests[n].from[2])

		assert.NoError(t, err)
		assert.InDeltaf(t, tests[n].to[0], h, precision, "h is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[1], s, precision, "s is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[2], l, precision, "l is wrong for test #%v", n+1)
	}
}

func TestRGBtoHSV_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-1, 0, 0},
		{2, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.RGBtoHSL(tests[n][0], tests[n][1], tests[n][2])
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestRGBtoHSV(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{0.00, 0.00, 0.00}, to: []float64{0.000, 0.00, 0.00}},
		{from: []float64{1.00, 1.00, 1.00}, to: []float64{0.000, 0.00, 1.00}},
		{from: []float64{1.00, 0.00, 0.00}, to: []float64{0.000, 1.00, 1.00}},
		{from: []float64{0.00, 1.00, 0.00}, to: []float64{120.0, 1.00, 1.00}},
		{from: []float64{0.00, 0.00, 1.00}, to: []float64{240.0, 1.00, 1.00}},
		{from: []float64{1.00, 1.00, 0.00}, to: []float64{60.00, 1.00, 1.00}},
		{from: []float64{1.00, 0.00, 1.00}, to: []float64{300.0, 1.00, 1.00}},
		{from: []float64{0.00, 1.00, 1.00}, to: []float64{180.0, 1.00, 1.00}},
		{from: []float64{1.00, 0.50, 0.00}, to: []float64{30.00, 1.00, 1.00}},
		{from: []float64{1.00, 0.50, 0.50}, to: []float64{0.000, 0.50, 1.00}},
		{from: []float64{0.50, 1.00, 0.00}, to: []float64{90.00, 1.00, 1.00}},
		{from: []float64{0.50, 1.00, 0.50}, to: []float64{120.0, 0.50, 1.00}},
		{from: []float64{0.50, 0.00, 1.00}, to: []float64{270.0, 1.00, 1.00}},
		{from: []float64{0.50, 0.50, 1.00}, to: []float64{240.0, 0.50, 1.00}},
		{from: []float64{0.50, 0.50, 0.50}, to: []float64{0.000, 0.00, 0.50}},
		{from: []float64{0.50, 0.50, 0.00}, to: []float64{60.00, 1.00, 0.50}},
		{from: []float64{0.50, 0.00, 0.50}, to: []float64{300.0, 1.00, 0.50}},
		{from: []float64{0.00, 0.50, 0.50}, to: []float64{180.0, 1.00, 0.50}},
		{from: []float64{0.50, 0.00, 0.00}, to: []float64{0.000, 1.00, 0.50}},
		{from: []float64{0.00, 0.50, 0.00}, to: []float64{120.0, 1.00, 0.50}},
		{from: []float64{0.00, 0.00, 0.50}, to: []float64{240.0, 1.00, 0.50}},
		{from: []float64{1.00, 0.50, 0.25}, to: []float64{20.00, 0.75, 1.00}},
		{from: []float64{0.50, 1.00, 0.25}, to: []float64{100.0, 0.75, 1.00}},
		{from: []float64{0.50, 0.25, 1.00}, to: []float64{260.0, 0.75, 1.00}},
		{from: []float64{0.25, 1.00, 0.50}, to: []float64{140.0, 0.75, 1.00}},
		{from: []float64{0.25, 0.50, 1.00}, to: []float64{220.0, 0.75, 1.00}},
		{from: []float64{0.75, 0.00, 0.00}, to: []float64{0.000, 1.00, 0.75}},
		{from: []float64{0.00, 0.75, 0.00}, to: []float64{120.0, 1.00, 0.75}},
		{from: []float64{0.00, 0.00, 0.75}, to: []float64{240.0, 1.00, 0.75}},
		{from: []float64{0.75, 0.75, 0.00}, to: []float64{60.00, 1.00, 0.75}},
		{from: []float64{0.75, 0.00, 0.75}, to: []float64{300.0, 1.00, 0.75}},
		{from: []float64{0.00, 0.75, 0.75}, to: []float64{180.0, 1.00, 0.75}},
		{from: []float64{0.75, 0.50, 0.00}, to: []float64{40.00, 1.00, 0.75}},
		{from: []float64{0.75, 0.00, 0.50}, to: []float64{320.0, 1.00, 0.75}},
		{from: []float64{0.75, 0.50, 0.50}, to: []float64{0.000, 1.0 / 3.0, 0.75}},
		{from: []float64{0.50, 0.75, 0.00}, to: []float64{80.00, 1.00, 0.75}},
		{from: []float64{0.00, 0.75, 0.50}, to: []float64{160.0, 1.00, 0.75}},
		{from: []float64{0.50, 0.75, 0.50}, to: []float64{120.0, 1.0 / 3.0, 0.75}},
		{from: []float64{0.50, 0.00, 0.75}, to: []float64{280.0, 1.00, 0.75}},
		{from: []float64{0.00, 0.50, 0.75}, to: []float64{200.0, 1.00, 0.75}},
		{from: []float64{0.50, 0.50, 0.75}, to: []float64{240.0, 1.0 / 3.0, 0.75}},
	}

	for n := 0; n < len(tests); n++ {
		h, s, v, err := gocolor.RGBtoHSV(tests[n].from[0], tests[n].from[1], tests[n].from[2])

		assert.NoError(t, err)
		assert.InDeltaf(t, tests[n].to[0], h, precision, "h is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[1], s, precision, "s is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[2], v, precision, "v is wrong for test #%v", n+1)
	}
}

func TestRGBtoYIQ_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-1, 0, 0},
		{2, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.RGBtoYIQ(tests[n][0], tests[n][1], tests[n][2])
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestRGBtoYIQ(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{0.00, 0.00, 0.00}, to: []float64{0.00000000, 00.00000000, 00.00000000}},
		{from: []float64{1.00, 1.00, 1.00}, to: []float64{1.00000000, 00.00000000, -0.00000001}},
		{from: []float64{1.00, 0.00, 0.00}, to: []float64{0.29895808, 00.59590296, 00.21133576}},
		{from: []float64{0.00, 1.00, 0.00}, to: []float64{0.58660979, -0.27405705, -0.52263517}},
		{from: []float64{0.00, 0.00, 1.00}, to: []float64{0.11443213, -0.32184591, 00.31129940}},
		{from: []float64{1.00, 1.00, 0.00}, to: []float64{0.88556787, 00.32184591, -0.31129941}},
		{from: []float64{1.00, 0.00, 1.00}, to: []float64{0.41339021, 00.27405705, 00.52263516}},
		{from: []float64{0.00, 1.00, 1.00}, to: []float64{0.70104192, -0.59590296, -0.21133577}},
		{from: []float64{1.00, 0.50, 0.00}, to: []float64{0.59226298, 00.45887444, -0.04998183}},
		{from: []float64{1.00, 0.50, 0.50}, to: []float64{0.64947904, 00.29795148, 00.10566788}},
		{from: []float64{0.50, 1.00, 0.00}, to: []float64{0.73608883, 00.02389443, -0.41696729}},
		{from: []float64{0.50, 1.00, 0.50}, to: []float64{0.79330490, -0.13702853, -0.26131759}},
		{from: []float64{0.50, 0.00, 1.00}, to: []float64{0.26391117, -0.02389443, 00.41696728}},
		{from: []float64{0.50, 0.50, 1.00}, to: []float64{0.55721607, -0.16092296, 00.15564970}},
		{from: []float64{0.50, 0.50, 0.50}, to: []float64{0.50000000, 00.00000000, 00.00000000}},
		{from: []float64{0.50, 0.50, 0.00}, to: []float64{0.44278394, 00.16092296, -0.15564971}},
		{from: []float64{0.50, 0.00, 0.50}, to: []float64{0.20669511, 00.13702853, 00.26131758}},
		{from: []float64{0.00, 0.50, 0.50}, to: []float64{0.35052096, -0.29795148, -0.10566789}},
		{from: []float64{0.50, 0.00, 0.00}, to: []float64{0.14947904, 00.29795148, 00.10566788}},
		{from: []float64{0.00, 0.50, 0.00}, to: []float64{0.29330490, -0.13702853, -0.26131759}},
		{from: []float64{0.00, 0.00, 0.50}, to: []float64{0.05721607, -0.16092296, 00.15564970}},
		{from: []float64{1.00, 0.50, 0.25}, to: []float64{0.62087101, 00.37841296, 00.02784303}},
		{from: []float64{0.50, 1.00, 0.25}, to: []float64{0.76469686, -0.05656705, -0.33914244}},
		{from: []float64{0.50, 0.25, 1.00}, to: []float64{0.41056362, -0.09240869, 00.28630849}},
		{from: []float64{0.25, 1.00, 0.50}, to: []float64{0.71856538, -0.28600427, -0.31415153}},
		{from: []float64{0.25, 0.50, 1.00}, to: []float64{0.48247655, -0.30989870, 00.10281576}},
		{from: []float64{0.75, 0.00, 0.00}, to: []float64{0.22421856, 00.44692722, 00.15850182}},
		{from: []float64{0.00, 0.75, 0.00}, to: []float64{0.43995734, -0.20554279, -0.39197638}},
		{from: []float64{0.00, 0.00, 0.75}, to: []float64{0.08582410, -0.24138443, 00.23347455}},
		{from: []float64{0.75, 0.75, 0.00}, to: []float64{0.66417590, 00.24138443, -0.23347456}},
		{from: []float64{0.75, 0.00, 0.75}, to: []float64{0.31004266, 00.20554279, 00.39197637}},
		{from: []float64{0.00, 0.75, 0.75}, to: []float64{0.52578144, -0.44692722, -0.15850183}},
		{from: []float64{0.75, 0.50, 0.00}, to: []float64{0.51752346, 00.30989870, -0.10281577}},
		{from: []float64{0.75, 0.00, 0.50}, to: []float64{0.28143462, 00.28600427, 00.31415152}},
		{from: []float64{0.75, 0.50, 0.50}, to: []float64{0.57473952, 00.14897574, 00.05283394}},
		{from: []float64{0.50, 0.75, 0.00}, to: []float64{0.58943638, 00.09240869, -0.28630850}},
		{from: []float64{0.00, 0.75, 0.50}, to: []float64{0.49717341, -0.36646574, -0.23632668}},
		{from: []float64{0.50, 0.75, 0.50}, to: []float64{0.64665245, -0.06851426, -0.13065880}},
		{from: []float64{0.50, 0.00, 0.75}, to: []float64{0.23530314, 00.05656705, 00.33914243}},
		{from: []float64{0.00, 0.50, 0.75}, to: []float64{0.37912899, -0.37841296, -0.02784304}},
		{from: []float64{0.50, 0.50, 0.75}, to: []float64{0.52860803, -0.08046148, 00.07782485}},
	}

	for n := 0; n < len(tests); n++ {
		y, i, q, err := gocolor.RGBtoYIQ(tests[n].from[0], tests[n].from[1], tests[n].from[2])

		assert.NoError(t, err)
		assert.InDeltaf(t, tests[n].to[0], y, precision, "y is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[1], i, precision, "i is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[2], q, precision, "q is wrong for test #%v", n+1)
	}
}

func TestRGBtoYUV(t *testing.T) {
	t.Run("convert RGB color to SDTV YUV", TestRGBtoSDYUV)
	t.Run("convert RGB color to HDTV YUV", TestRGBtoHDYUV)
}

func TestRGBtoSDYUV_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-1, 0, 0},
		{2, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.RGBtoSDYUV(tests[n][0], tests[n][1], tests[n][2])
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestRGBtoSDYUV(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{0.00, 0.00, 0.00}, to: []float64{0.00000000, 00.00000000, 00.00000000}},
		{from: []float64{1.00, 1.00, 1.00}, to: []float64{1.00000000, 00.00001000, 00.00000000}},
		{from: []float64{1.00, 0.00, 0.00}, to: []float64{0.29900000, -0.14713000, 00.61500000}},
		{from: []float64{0.00, 1.00, 0.00}, to: []float64{0.58700000, -0.28886000, -0.51499000}},
		{from: []float64{0.00, 0.00, 1.00}, to: []float64{0.11400000, 00.43600000, -0.10001000}},
		{from: []float64{1.00, 1.00, 0.00}, to: []float64{0.88600000, -0.43599000, 00.10001000}},
		{from: []float64{1.00, 0.00, 1.00}, to: []float64{0.41300000, 00.28887000, 00.51499000}},
		{from: []float64{0.00, 1.00, 1.00}, to: []float64{0.70100000, 00.14714000, -0.61500000}},
	}

	for n := 0; n < len(tests); n++ {
		y, u, v, err := gocolor.RGBtoSDYUV(tests[n].from[0], tests[n].from[1], tests[n].from[2])

		assert.NoError(t, err)
		assert.InDeltaf(t, tests[n].to[0], y, precision, "y is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[1], u, precision, "u is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[2], v, precision, "v is wrong for test #%v", n+1)
	}
}

func TestRGBtoHDYUV_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-1, 0, 0},
		{2, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.RGBtoHDYUV(tests[n][0], tests[n][1], tests[n][2])
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestRGBtoHDYUV(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{0.00, 0.00, 0.00}, to: []float64{0.00000000, 00.00000000, 00.00000000}},
		{from: []float64{1.00, 1.00, 1.00}, to: []float64{1.00000000, 00.00000000, 00.11278000}},
		{from: []float64{1.00, 0.00, 0.00}, to: []float64{0.21260000, -0.09991000, 00.61500000}},
		{from: []float64{0.00, 1.00, 0.00}, to: []float64{0.71520000, -0.33609000, -0.55861000}},
		{from: []float64{0.00, 0.00, 1.00}, to: []float64{0.07220000, 00.43600000, 00.05639000}},
		{from: []float64{1.00, 1.00, 0.00}, to: []float64{0.92780000, -0.43600000, 00.05639000}},
		{from: []float64{1.00, 0.00, 1.00}, to: []float64{0.28480000, 00.33609000, 00.67139000}},
		{from: []float64{0.00, 1.00, 1.00}, to: []float64{0.78740000, 00.09991000, -0.50222000}},
	}

	for n := 0; n < len(tests); n++ {
		y, u, v, err := gocolor.RGBtoHDYUV(tests[n].from[0], tests[n].from[1], tests[n].from[2])

		assert.NoError(t, err)
		assert.InDeltaf(t, tests[n].to[0], y, precision, "y is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[1], u, precision, "u is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[2], v, precision, "v is wrong for test #%v", n+1)
	}
}

func TestRGBtoCMY_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-1, 0, 0},
		{2, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.RGBtoCMY(tests[n][0], tests[n][1], tests[n][2])
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestRGBtoCMY(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{0.00, 0.00, 0.00}, to: []float64{1.00, 1.00, 1.00}},
		{from: []float64{1.00, 1.00, 1.00}, to: []float64{0.00, 0.00, 0.00}},
		{from: []float64{1.00, 0.00, 0.00}, to: []float64{0.00, 1.00, 1.00}},
		{from: []float64{0.00, 1.00, 0.00}, to: []float64{1.00, 0.00, 1.00}},
		{from: []float64{0.00, 0.00, 1.00}, to: []float64{1.00, 1.00, 0.00}},
		{from: []float64{1.00, 1.00, 0.00}, to: []float64{0.00, 0.00, 1.00}},
		{from: []float64{1.00, 0.00, 1.00}, to: []float64{0.00, 1.00, 0.00}},
		{from: []float64{0.00, 1.00, 1.00}, to: []float64{1.00, 0.00, 0.00}},
		{from: []float64{1.00, 0.50, 0.00}, to: []float64{0.00, 0.50, 1.00}},
		{from: []float64{1.00, 0.50, 0.50}, to: []float64{0.00, 0.50, 0.50}},
		{from: []float64{0.50, 1.00, 0.00}, to: []float64{0.50, 0.00, 1.00}},
		{from: []float64{0.50, 1.00, 0.50}, to: []float64{0.50, 0.00, 0.50}},
		{from: []float64{0.50, 0.00, 1.00}, to: []float64{0.50, 1.00, 0.00}},
		{from: []float64{0.50, 0.50, 1.00}, to: []float64{0.50, 0.50, 0.00}},
		{from: []float64{0.50, 0.50, 0.50}, to: []float64{0.50, 0.50, 0.50}},
		{from: []float64{0.50, 0.50, 0.00}, to: []float64{0.50, 0.50, 1.00}},
		{from: []float64{0.50, 0.00, 0.50}, to: []float64{0.50, 1.00, 0.50}},
		{from: []float64{0.00, 0.50, 0.50}, to: []float64{1.00, 0.50, 0.50}},
		{from: []float64{0.50, 0.00, 0.00}, to: []float64{0.50, 1.00, 1.00}},
		{from: []float64{0.00, 0.50, 0.00}, to: []float64{1.00, 0.50, 1.00}},
		{from: []float64{0.00, 0.00, 0.50}, to: []float64{1.00, 1.00, 0.50}},
		{from: []float64{1.00, 0.50, 0.25}, to: []float64{0.00, 0.50, 0.75}},
		{from: []float64{0.50, 1.00, 0.25}, to: []float64{0.50, 0.00, 0.75}},
		{from: []float64{0.50, 0.25, 1.00}, to: []float64{0.50, 0.75, 0.00}},
		{from: []float64{0.25, 1.00, 0.50}, to: []float64{0.75, 0.00, 0.50}},
		{from: []float64{0.25, 0.50, 1.00}, to: []float64{0.75, 0.50, 0.00}},
		{from: []float64{0.75, 0.00, 0.00}, to: []float64{0.25, 1.00, 1.00}},
		{from: []float64{0.00, 0.75, 0.00}, to: []float64{1.00, 0.25, 1.00}},
		{from: []float64{0.00, 0.00, 0.75}, to: []float64{1.00, 1.00, 0.25}},
		{from: []float64{0.75, 0.75, 0.00}, to: []float64{0.25, 0.25, 1.00}},
		{from: []float64{0.75, 0.00, 0.75}, to: []float64{0.25, 1.00, 0.25}},
		{from: []float64{0.00, 0.75, 0.75}, to: []float64{1.00, 0.25, 0.25}},
		{from: []float64{0.75, 0.50, 0.00}, to: []float64{0.25, 0.50, 1.00}},
		{from: []float64{0.75, 0.00, 0.50}, to: []float64{0.25, 1.00, 0.50}},
		{from: []float64{0.75, 0.50, 0.50}, to: []float64{0.25, 0.50, 0.50}},
		{from: []float64{0.50, 0.75, 0.00}, to: []float64{0.50, 0.25, 1.00}},
		{from: []float64{0.00, 0.75, 0.50}, to: []float64{1.00, 0.25, 0.50}},
		{from: []float64{0.50, 0.75, 0.50}, to: []float64{0.50, 0.25, 0.50}},
		{from: []float64{0.50, 0.00, 0.75}, to: []float64{0.50, 1.00, 0.25}},
		{from: []float64{0.00, 0.50, 0.75}, to: []float64{1.00, 0.50, 0.25}},
		{from: []float64{0.50, 0.50, 0.75}, to: []float64{0.50, 0.50, 0.25}},
	}

	for n := 0; n < len(tests); n++ {
		c, m, y, err := gocolor.RGBtoCMY(tests[n].from[0], tests[n].from[1], tests[n].from[2])

		assert.NoError(t, err)
		assert.InDeltaf(t, tests[n].to[0], c, precision, "c is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[1], m, precision, "m is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[2], y, precision, "y is wrong for test #%v", n+1)
	}
}

func TestRGBtoHEX_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-1, 0, 0},
		{2, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, err := gocolor.RGBtoHEX(tests[n][0], tests[n][1], tests[n][2])
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestRGBtoHEX(t *testing.T) {
	tests := []struct {
		from []float64
		to   string
	}{
		{from: []float64{0.00, 0.00, 0.00}, to: "#000000"},
		{from: []float64{1.00, 1.00, 1.00}, to: "#FFFFFF"},
		{from: []float64{1.00, 0.00, 0.00}, to: "#FF0000"},
		{from: []float64{0.00, 1.00, 0.00}, to: "#00FF00"},
		{from: []float64{0.00, 0.00, 1.00}, to: "#0000FF"},
		{from: []float64{1.00, 1.00, 0.00}, to: "#FFFF00"},
		{from: []float64{1.00, 0.00, 1.00}, to: "#FF00FF"},
		{from: []float64{0.00, 1.00, 1.00}, to: "#00FFFF"},
		{from: []float64{1.00, 0.50, 0.00}, to: "#FF8000"},
		{from: []float64{1.00, 0.50, 0.50}, to: "#FF8080"},
		{from: []float64{0.50, 1.00, 0.00}, to: "#80FF00"},
		{from: []float64{0.50, 1.00, 0.50}, to: "#80FF80"},
		{from: []float64{0.50, 0.00, 1.00}, to: "#8000FF"},
		{from: []float64{0.50, 0.50, 1.00}, to: "#8080FF"},
		{from: []float64{0.50, 0.50, 0.50}, to: "#808080"},
		{from: []float64{0.50, 0.50, 0.00}, to: "#808000"},
		{from: []float64{0.50, 0.00, 0.50}, to: "#800080"},
		{from: []float64{0.00, 0.50, 0.50}, to: "#008080"},
		{from: []float64{0.50, 0.00, 0.00}, to: "#800000"},
		{from: []float64{0.00, 0.50, 0.00}, to: "#008000"},
		{from: []float64{0.00, 0.00, 0.50}, to: "#000080"},
		{from: []float64{1.00, 0.50, 0.25}, to: "#FF8040"},
		{from: []float64{0.50, 1.00, 0.25}, to: "#80FF40"},
		{from: []float64{0.50, 0.25, 1.00}, to: "#8040FF"},
		{from: []float64{0.25, 1.00, 0.50}, to: "#40FF80"},
		{from: []float64{0.25, 0.50, 1.00}, to: "#4080FF"},
		{from: []float64{0.75, 0.00, 0.00}, to: "#BF0000"},
		{from: []float64{0.00, 0.75, 0.00}, to: "#00BF00"},
		{from: []float64{0.00, 0.00, 0.75}, to: "#0000BF"},
		{from: []float64{0.75, 0.75, 0.00}, to: "#BFBF00"},
		{from: []float64{0.75, 0.00, 0.75}, to: "#BF00BF"},
		{from: []float64{0.00, 0.75, 0.75}, to: "#00BFBF"},
		{from: []float64{0.75, 0.50, 0.00}, to: "#BF8000"},
		{from: []float64{0.75, 0.00, 0.50}, to: "#BF0080"},
		{from: []float64{0.75, 0.50, 0.50}, to: "#BF8080"},
		{from: []float64{0.50, 0.75, 0.00}, to: "#80BF00"},
		{from: []float64{0.00, 0.75, 0.50}, to: "#00BF80"},
		{from: []float64{0.50, 0.75, 0.50}, to: "#80BF80"},
		{from: []float64{0.50, 0.00, 0.75}, to: "#8000BF"},
		{from: []float64{0.00, 0.50, 0.75}, to: "#0080BF"},
		{from: []float64{0.50, 0.50, 0.75}, to: "#8080BF"},
	}

	for n := 0; n < len(tests); n++ {
		hex, err := gocolor.RGBtoHEX(tests[n].from[0], tests[n].from[1], tests[n].from[2])
		assert.NoError(t, err)
		assert.Equal(t, tests[n].to, hex)
	}
}

func TestRGBtoXYZ_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-1, 0, 0},
		{2, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.RGBtoXYZ(tests[n][0], tests[n][1], tests[n][2], gocolor.AdobeRGB)
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}

	_, _, _, err := gocolor.RGBtoXYZ(0, 0, 0, "invalid space")
	assert.Errorf(t, err, "invalid color space should return an error")
}

func TestRGBtoXYZ(t *testing.T) {
	t.Run("convert sRGB color to XYZ", TestSRGBtoXYZ)
	t.Run("convert ITU-R BT.2020 RGB color to XYZ", TestBT2020toXYZ)
	t.Run("convert 12bit ITU-R BT.2020 RGB color to XYZ", TestBT202012btoXYZ)
	t.Run("convert Adobe RGB color to XYZ", TestAdobeRGBtoXYZ)
}

func TestSRGBtoXYZ(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{0.00, 0.00, 0.00}, to: []float64{0.00000000, 0.00000000, 0.00000000}},
		{from: []float64{1.00, 1.00, 1.00}, to: []float64{0.95047000, 1.00000010, 1.08883000}},
		{from: []float64{1.00, 0.00, 0.00}, to: []float64{0.41245640, 0.21267290, 0.01933390}},
		{from: []float64{0.00, 1.00, 0.00}, to: []float64{0.35757610, 0.71515220, 0.11919200}},
		{from: []float64{0.00, 0.00, 1.00}, to: []float64{0.18043750, 0.07217500, 0.95030410}},
		{from: []float64{1.00, 1.00, 0.00}, to: []float64{0.77003250, 0.92782510, 0.13852590}},
		{from: []float64{1.00, 0.00, 1.00}, to: []float64{0.59289390, 0.28484790, 0.96963800}},
		{from: []float64{0.00, 1.00, 1.00}, to: []float64{0.53801360, 0.78732720, 1.06949610}},
		{from: []float64{1.00, 0.50, 0.00}, to: []float64{0.48899240, 0.36574489, 0.04484589}},
		{from: []float64{1.00, 0.50, 0.50}, to: []float64{0.52761345, 0.38119331, 0.24825007}},
		{from: []float64{0.50, 1.00, 0.00}, to: []float64{0.44585874, 0.76067295, 0.12333025}},
		{from: []float64{0.50, 1.00, 0.50}, to: []float64{0.48447979, 0.77612137, 0.32673442}},
		{from: []float64{0.50, 0.00, 1.00}, to: []float64{0.26872013, 0.11769575, 0.95444235}},
		{from: []float64{0.50, 0.50, 1.00}, to: []float64{0.34525614, 0.27076774, 0.97995434}},
		{from: []float64{0.50, 0.50, 0.50}, to: []float64{0.20343968, 0.21404116, 0.23305442}},
		{from: []float64{0.50, 0.50, 0.00}, to: []float64{0.16481864, 0.19859274, 0.02965024}},
		{from: []float64{0.50, 0.00, 0.50}, to: []float64{0.12690369, 0.06096917, 0.20754242}},
		{from: []float64{0.00, 0.50, 0.50}, to: []float64{0.11515705, 0.16852041, 0.22891617}},
		{from: []float64{0.50, 0.00, 0.00}, to: []float64{0.08828264, 0.04552075, 0.00413825}},
		{from: []float64{0.00, 0.50, 0.00}, to: []float64{0.07653600, 0.15307199, 0.02551199}},
		{from: []float64{0.00, 0.00, 0.50}, to: []float64{0.03862105, 0.01544842, 0.20340417}},
		{from: []float64{1.00, 0.50, 0.25}, to: []float64{0.49817235, 0.36941687, 0.09319365}},
		{from: []float64{0.50, 1.00, 0.25}, to: []float64{0.45503869, 0.76434493, 0.17167801}},
		{from: []float64{0.50, 0.25, 1.00}, to: []float64{0.28691221, 0.15407990, 0.96050637}},
		{from: []float64{0.25, 1.00, 0.50}, to: []float64{0.41718132, 0.74142059, 0.32357981}},
		{from: []float64{0.25, 0.50, 1.00}, to: []float64{0.27795766, 0.23606696, 0.97679973}},
		{from: []float64{0.75, 0.00, 0.00}, to: []float64{0.21551736, 0.11112617, 0.01010238}},
		{from: []float64{0.00, 0.75, 0.00}, to: []float64{0.18684122, 0.37368244, 0.06228039}},
		{from: []float64{0.00, 0.00, 0.75}, to: []float64{0.09428248, 0.03771299, 0.49655438}},
		{from: []float64{0.75, 0.75, 0.00}, to: []float64{0.40235858, 0.48480861, 0.07238277}},
		{from: []float64{0.75, 0.00, 0.75}, to: []float64{0.30979984, 0.14883917, 0.50665676}},
		{from: []float64{0.00, 0.75, 0.75}, to: []float64{0.28112370, 0.41139543, 0.55883476}},
		{from: []float64{0.75, 0.50, 0.00}, to: []float64{0.29205336, 0.26419817, 0.03561437}},
		{from: []float64{0.75, 0.00, 0.50}, to: []float64{0.25413841, 0.12657459, 0.21350655}},
		{from: []float64{0.75, 0.50, 0.50}, to: []float64{0.33067440, 0.27964659, 0.23901855}},
		{from: []float64{0.50, 0.75, 0.00}, to: []float64{0.27512386, 0.41920319, 0.06641864}},
		{from: []float64{0.00, 0.75, 0.50}, to: []float64{0.22546227, 0.38913086, 0.26568456}},
		{from: []float64{0.50, 0.75, 0.50}, to: []float64{0.31374491, 0.43465161, 0.26982281}},
		{from: []float64{0.50, 0.00, 0.75}, to: []float64{0.18256512, 0.08323374, 0.50069263}},
		{from: []float64{0.00, 0.50, 0.75}, to: []float64{0.17081848, 0.19078499, 0.52206637}},
		{from: []float64{0.50, 0.50, 0.75}, to: []float64{0.25910112, 0.23630574, 0.52620462}},
	}

	for n := 0; n < len(tests); n++ {
		x, y, z, err := gocolor.RGBtoXYZ(tests[n].from[0], tests[n].from[1], tests[n].from[2], gocolor.SRGB)

		assert.NoError(t, err)
		assert.InDeltaf(t, tests[n].to[0], x, precision, "x is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[1], y, precision, "y is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[2], z, precision, "z is wrong for test #%v", n+1)
	}
}

func TestBT2020toXYZ(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{0.00, 0.00, 0.00}, to: []float64{0.00000000, 0.00000000, 0.00000000}},
		{from: []float64{1.00, 1.00, 1.00}, to: []float64{0.95045593, 1.00000000, 1.08905775}},
		{from: []float64{1.00, 0.00, 0.00}, to: []float64{0.63695805, 0.26270021, 0.00000000}},
		{from: []float64{0.00, 1.00, 0.00}, to: []float64{0.14461690, 0.67799807, 0.02807269}},
		{from: []float64{0.00, 0.00, 1.00}, to: []float64{0.16888098, 0.05930172, 1.06098506}},
		{from: []float64{1.00, 1.00, 0.00}, to: []float64{0.78157495, 0.94069828, 0.02807269}},
		{from: []float64{1.00, 0.00, 1.00}, to: []float64{0.80583902, 0.32200193, 1.06098506}},
		{from: []float64{0.00, 1.00, 1.00}, to: []float64{0.31349788, 0.73729979, 1.08905775}},
		{from: []float64{1.00, 0.50, 0.00}, to: []float64{0.67449906, 0.43870132, 0.00728737}},
		{from: []float64{1.00, 0.50, 0.50}, to: []float64{0.71833877, 0.45409542, 0.28270785}},
		{from: []float64{0.50, 1.00, 0.00}, to: []float64{0.30996446, 0.74619226, 0.02807269}},
		{from: []float64{0.50, 1.00, 0.50}, to: []float64{0.35380417, 0.76158636, 0.30349317}},
		{from: []float64{0.50, 0.00, 1.00}, to: []float64{0.33422853, 0.12749591, 1.06098506}},
		{from: []float64{0.50, 0.50, 1.00}, to: []float64{0.37176955, 0.30349702, 1.06827243}},
		{from: []float64{0.50, 0.50, 0.50}, to: []float64{0.24672828, 0.25958940, 0.28270785}},
		{from: []float64{0.50, 0.50, 0.00}, to: []float64{0.20288857, 0.24419530, 0.00728737}},
		{from: []float64{0.50, 0.00, 0.50}, to: []float64{0.20918727, 0.08358829, 0.27542048}},
		{from: []float64{0.00, 0.50, 0.50}, to: []float64{0.08138073, 0.19139521, 0.28270785}},
		{from: []float64{0.50, 0.00, 0.00}, to: []float64{0.16534756, 0.06819419, 0.00000000}},
		{from: []float64{0.00, 0.50, 0.00}, to: []float64{0.03754102, 0.17600111, 0.00728737}},
		{from: []float64{0.00, 0.00, 0.50}, to: []float64{0.04383971, 0.01539410, 0.27542048}},
		{from: []float64{1.00, 0.50, 0.25}, to: []float64{0.68769777, 0.44333598, 0.09020747}},
		{from: []float64{0.50, 1.00, 0.25}, to: []float64{0.32316316, 0.75082692, 0.11099279}},
		{from: []float64{0.50, 0.25, 1.00}, to: []float64{0.34553090, 0.18048408, 1.06317905}},
		{from: []float64{0.25, 1.00, 0.50}, to: []float64{0.23823735, 0.71392321, 0.30349317}},
		{from: []float64{0.25, 0.50, 1.00}, to: []float64{0.25620273, 0.25583387, 1.06827243}},
		{from: []float64{0.75, 0.00, 0.00}, to: []float64{0.35894006, 0.14803743, 0.00000000}},
		{from: []float64{0.00, 0.75, 0.00}, to: []float64{0.08149485, 0.38206703, 0.01581959}},
		{from: []float64{0.00, 0.00, 0.75}, to: []float64{0.09516820, 0.03341784, 0.59788874}},
		{from: []float64{0.75, 0.75, 0.00}, to: []float64{0.44043491, 0.53010446, 0.01581959}},
		{from: []float64{0.75, 0.00, 0.75}, to: []float64{0.45410826, 0.18145527, 0.59788874}},
		{from: []float64{0.00, 0.75, 0.75}, to: []float64{0.17666305, 0.41548487, 0.61370833}},
		{from: []float64{0.75, 0.50, 0.00}, to: []float64{0.39648108, 0.32403854, 0.00728737}},
		{from: []float64{0.75, 0.00, 0.50}, to: []float64{0.40277977, 0.16343152, 0.27542048}},
		{from: []float64{0.75, 0.50, 0.50}, to: []float64{0.44032079, 0.33943264, 0.28270785}},
		{from: []float64{0.50, 0.75, 0.00}, to: []float64{0.24684241, 0.45026122, 0.01581959}},
		{from: []float64{0.00, 0.75, 0.50}, to: []float64{0.12533456, 0.39746113, 0.29124006}},
		{from: []float64{0.50, 0.75, 0.50}, to: []float64{0.29068212, 0.46565532, 0.29124006}},
		{from: []float64{0.50, 0.00, 0.75}, to: []float64{0.26051575, 0.10161203, 0.59788874}},
		{from: []float64{0.00, 0.50, 0.75}, to: []float64{0.13270921, 0.20941895, 0.60517611}},
		{from: []float64{0.50, 0.50, 0.75}, to: []float64{0.29805677, 0.27761314, 0.60517611}},
	}

	for n := 0; n < len(tests); n++ {
		x, y, z, err := gocolor.RGBtoXYZ(tests[n].from[0], tests[n].from[1], tests[n].from[2], gocolor.BT2020)

		assert.NoError(t, err)
		assert.InDeltaf(t, tests[n].to[0], x, precision, "x is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[1], y, precision, "y is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[2], z, precision, "z is wrong for test #%v", n+1)
	}
}

func TestBT202012btoXYZ(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{0.00, 0.00, 0.00}, to: []float64{0.00000000, 0.00000000, 0.00000000}},
		{from: []float64{1.00, 1.00, 1.00}, to: []float64{0.95045593, 1.00000000, 1.08905775}},
		{from: []float64{1.00, 0.00, 0.00}, to: []float64{0.63695805, 0.26270021, 0.00000000}},
		{from: []float64{0.00, 1.00, 0.00}, to: []float64{0.14461690, 0.67799807, 0.02807269}},
		{from: []float64{0.00, 0.00, 1.00}, to: []float64{0.16888098, 0.05930172, 1.06098506}},
		{from: []float64{1.00, 1.00, 0.00}, to: []float64{0.78157495, 0.94069828, 0.02807269}},
		{from: []float64{1.00, 0.00, 1.00}, to: []float64{0.80583902, 0.32200193, 1.06098506}},
		{from: []float64{0.00, 1.00, 1.00}, to: []float64{0.31349788, 0.73729979, 1.08905775}},
		{from: []float64{1.00, 0.50, 0.00}, to: []float64{0.67451807, 0.43879043, 0.00729106}},
		{from: []float64{1.00, 0.50, 0.50}, to: []float64{0.71837998, 0.45419232, 0.28285098}},
		{from: []float64{0.50, 1.00, 0.00}, to: []float64{0.31004817, 0.74622679, 0.02807269}},
		{from: []float64{0.50, 1.00, 0.50}, to: []float64{0.35391008, 0.76162868, 0.30363261}},
		{from: []float64{0.50, 0.00, 1.00}, to: []float64{0.33431225, 0.12753043, 1.06098506}},
		{from: []float64{0.50, 0.50, 1.00}, to: []float64{0.37187227, 0.30362065, 1.06827612}},
		{from: []float64{0.50, 0.50, 0.50}, to: []float64{0.24685320, 0.25972083, 0.28285098}},
		{from: []float64{0.50, 0.50, 0.00}, to: []float64{0.20299129, 0.24431894, 0.00729106}},
		{from: []float64{0.50, 0.00, 0.50}, to: []float64{0.20929318, 0.08363061, 0.27555992}},
		{from: []float64{0.00, 0.50, 0.50}, to: []float64{0.08142193, 0.19149211, 0.28285098}},
		{from: []float64{0.50, 0.00, 0.00}, to: []float64{0.16543127, 0.06822872, 0.00000000}},
		{from: []float64{0.00, 0.50, 0.00}, to: []float64{0.03756002, 0.17609022, 0.00729106}},
		{from: []float64{0.00, 0.00, 0.50}, to: []float64{0.04386191, 0.01540189, 0.27555992}},
		{from: []float64{1.00, 0.50, 0.25}, to: []float64{0.68773398, 0.44343113, 0.09031926}},
		{from: []float64{0.50, 1.00, 0.25}, to: []float64{0.32326408, 0.75086749, 0.11110089}},
		{from: []float64{0.50, 0.25, 1.00}, to: []float64{0.34562935, 0.18058769, 1.06318191}},
		{from: []float64{0.25, 1.00, 0.50}, to: []float64{0.23832445, 0.71395777, 0.30363261}},
		{from: []float64{0.25, 0.50, 1.00}, to: []float64{0.25628664, 0.25594974, 1.06827612}},
		{from: []float64{0.75, 0.00, 0.00}, to: []float64{0.35900417, 0.14806386, 0.00000000}},
		{from: []float64{0.00, 0.75, 0.00}, to: []float64{0.08150940, 0.38213526, 0.01582241}},
		{from: []float64{0.00, 0.00, 0.75}, to: []float64{0.09518519, 0.03342381, 0.59799551}},
		{from: []float64{0.75, 0.75, 0.00}, to: []float64{0.44051357, 0.53019913, 0.01582241}},
		{from: []float64{0.75, 0.00, 0.75}, to: []float64{0.45418936, 0.18148767, 0.59799551}},
		{from: []float64{0.00, 0.75, 0.75}, to: []float64{0.17669459, 0.41555907, 0.61381793}},
		{from: []float64{0.75, 0.50, 0.00}, to: []float64{0.39656419, 0.32415408, 0.00729106}},
		{from: []float64{0.75, 0.00, 0.50}, to: []float64{0.40286607, 0.16346576, 0.27555992}},
		{from: []float64{0.75, 0.50, 0.50}, to: []float64{0.44042609, 0.33955598, 0.28285098}},
		{from: []float64{0.50, 0.75, 0.00}, to: []float64{0.24694067, 0.45036398, 0.01582241}},
		{from: []float64{0.00, 0.75, 0.50}, to: []float64{0.12537131, 0.39753715, 0.29138233}},
		{from: []float64{0.50, 0.75, 0.50}, to: []float64{0.29080258, 0.46576587, 0.29138233}},
		{from: []float64{0.50, 0.00, 0.75}, to: []float64{0.26061646, 0.10165252, 0.59799551}},
		{from: []float64{0.00, 0.50, 0.75}, to: []float64{0.13274521, 0.20951403, 0.60528658}},
		{from: []float64{0.50, 0.50, 0.75}, to: []float64{0.29817648, 0.27774274, 0.60528658}},
	}

	for n := 0; n < len(tests); n++ {
		x, y, z, err := gocolor.RGBtoXYZ(tests[n].from[0], tests[n].from[1], tests[n].from[2], gocolor.BT202012b)

		assert.NoError(t, err)
		assert.InDeltaf(t, tests[n].to[0], x, precision, "x is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[1], y, precision, "y is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[2], z, precision, "z is wrong for test #%v", n+1)
	}
}

func TestAdobeRGBtoXYZ(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{0.00, 0.00, 0.00}, to: []float64{0.00000000, 0.00000000, 0.00000000}},
		{from: []float64{1.00, 1.00, 1.00}, to: []float64{0.95047010, 1.00000010, 1.08883000}},
		{from: []float64{1.00, 0.00, 0.00}, to: []float64{0.57673090, 0.29737690, 0.02703430}},
		{from: []float64{0.00, 1.00, 0.00}, to: []float64{0.18555400, 0.62734910, 0.07068720}},
		{from: []float64{0.00, 0.00, 1.00}, to: []float64{0.18818520, 0.07527410, 0.99110850}},
		{from: []float64{1.00, 1.00, 0.00}, to: []float64{0.76228490, 0.92472600, 0.09772150}},
		{from: []float64{1.00, 0.00, 1.00}, to: []float64{0.76491610, 0.37265100, 1.01814280}},
		{from: []float64{0.00, 1.00, 1.00}, to: []float64{0.37373920, 0.70262320, 1.06179570}},
		{from: []float64{1.00, 0.50, 0.00}, to: []float64{0.61711443, 0.43391168, 0.04241850}},
		{from: []float64{1.00, 0.50, 0.50}, to: []float64{0.65807062, 0.45029416, 0.25812101}},
		{from: []float64{0.50, 1.00, 0.00}, to: []float64{0.31107235, 0.69206951, 0.07657088}},
		{from: []float64{0.50, 1.00, 0.50}, to: []float64{0.35202854, 0.70845198, 0.29227340}},
		{from: []float64{0.50, 0.00, 1.00}, to: []float64{0.31370355, 0.13999451, 0.99699218}},
		{from: []float64{0.50, 0.50, 1.00}, to: []float64{0.35408709, 0.27652929, 1.01237638}},
		{from: []float64{0.50, 0.50, 0.50}, to: []float64{0.20685807, 0.21763766, 0.23697039}},
		{from: []float64{0.50, 0.50, 0.00}, to: []float64{0.16590189, 0.20125519, 0.02126788}},
		{from: []float64{0.50, 0.00, 0.50}, to: []float64{0.16647454, 0.08110288, 0.22158620}},
		{from: []float64{0.00, 0.50, 0.50}, to: []float64{0.08133972, 0.15291726, 0.23108671}},
		{from: []float64{0.50, 0.00, 0.00}, to: []float64{0.12551835, 0.06472041, 0.00588368}},
		{from: []float64{0.00, 0.50, 0.00}, to: []float64{0.04038353, 0.13653478, 0.01538420}},
		{from: []float64{0.00, 0.00, 0.50}, to: []float64{0.04095618, 0.01638248, 0.21570252}},
		{from: []float64{1.00, 0.50, 0.25}, to: []float64{0.62602804, 0.43747712, 0.08936348}},
		{from: []float64{0.50, 1.00, 0.25}, to: []float64{0.31998596, 0.69563495, 0.12351587}},
		{from: []float64{0.50, 0.25, 1.00}, to: []float64{0.32249253, 0.16970961, 1.00034036}},
		{from: []float64{0.25, 1.00, 0.50}, to: []float64{0.25382770, 0.65781717, 0.28767023}},
		{from: []float64{0.25, 0.50, 1.00}, to: []float64{0.25588625, 0.22589447, 1.00777321}},
		{from: []float64{0.75, 0.00, 0.00}, to: []float64{0.30627250, 0.15792177, 0.01435654}},
		{from: []float64{0.00, 0.75, 0.00}, to: []float64{0.09853831, 0.33315325, 0.03753838}},
		{from: []float64{0.00, 0.00, 0.75}, to: []float64{0.09993560, 0.03997425, 0.52632740}},
		{from: []float64{0.75, 0.75, 0.00}, to: []float64{0.40481081, 0.49107503, 0.05189493}},
		{from: []float64{0.75, 0.00, 0.75}, to: []float64{0.40620810, 0.19789602, 0.54068394}},
		{from: []float64{0.00, 0.75, 0.75}, to: []float64{0.19847391, 0.37312751, 0.56386578}},
		{from: []float64{0.75, 0.50, 0.00}, to: []float64{0.34665603, 0.29445655, 0.02974074}},
		{from: []float64{0.75, 0.00, 0.50}, to: []float64{0.34722868, 0.17430425, 0.23005906}},
		{from: []float64{0.75, 0.50, 0.50}, to: []float64{0.38761222, 0.31083903, 0.24544326}},
		{from: []float64{0.50, 0.75, 0.00}, to: []float64{0.22405666, 0.39787366, 0.04342206}},
		{from: []float64{0.00, 0.75, 0.50}, to: []float64{0.13949449, 0.34953573, 0.25324090}},
		{from: []float64{0.50, 0.75, 0.50}, to: []float64{0.26501284, 0.41425614, 0.25912458}},
		{from: []float64{0.50, 0.00, 0.75}, to: []float64{0.22545396, 0.10469466, 0.53221108}},
		{from: []float64{0.00, 0.50, 0.75}, to: []float64{0.14031914, 0.17650903, 0.54171160}},
		{from: []float64{0.50, 0.50, 0.75}, to: []float64{0.26583749, 0.24122944, 0.54759528}},
	}

	for n := 0; n < len(tests); n++ {
		x, y, z, err := gocolor.RGBtoXYZ(tests[n].from[0], tests[n].from[1], tests[n].from[2], gocolor.AdobeRGB)

		assert.NoError(t, err)
		assert.InDeltaf(t, tests[n].to[0], x, precision, "x is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[1], y, precision, "y is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[2], z, precision, "z is wrong for test #%v", n+1)
	}
}

////////////////////////////////////////

func TestToRGB(t *testing.T) {
	t.Run("convert RGB color to HSL", TestRGBtoHSL)
	t.Run("convert HSL to RGB", TestHSLtoRGB)
	t.Run("convert HSV to RGB", TestHSVtoRGB)
	t.Run("convert YIQ to RGB", TestYIQtoRGB)
	t.Run("convert YUV to RGB", TestYUVtoRGB)
	t.Run("convert CMY to RGB", TestCMYtoRGB)
	t.Run("convert HEX to RGB", TestHEXtoRGB)
	t.Run("convert XYZ to RGB", TestXYZtoRGB)
}

func TestHSLtoRGB_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-10, 0, 0},
		{390, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.HSLtoRGB(tests[n][0], tests[n][1], tests[n][2])
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestHSLtoRGB(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{000.0, 0.0000, 0.0000}, to: []float64{0.00, 0.00, 0.00}},
		{from: []float64{000.0, 0.0000, 1.0000}, to: []float64{1.00, 1.00, 1.00}},
		{from: []float64{000.0, 1.0000, 0.5000}, to: []float64{1.00, 0.00, 0.00}},
		{from: []float64{030.0, 1.0000, 0.5000}, to: []float64{1.00, 0.50, 0.00}},
		{from: []float64{060.0, 1.0000, 0.5000}, to: []float64{1.00, 1.00, 0.00}},
		{from: []float64{090.0, 1.0000, 0.5000}, to: []float64{0.50, 1.00, 0.00}},
		{from: []float64{120.0, 1.0000, 0.5000}, to: []float64{0.00, 1.00, 0.00}},
		{from: []float64{150.0, 1.0000, 0.5000}, to: []float64{0.00, 1.00, 0.50}},
		{from: []float64{180.0, 1.0000, 0.5000}, to: []float64{0.00, 1.00, 1.00}},
		{from: []float64{210.0, 1.0000, 0.5000}, to: []float64{0.00, 0.50, 1.00}},
		{from: []float64{240.0, 1.0000, 0.5000}, to: []float64{0.00, 0.00, 1.00}},
		{from: []float64{270.0, 1.0000, 0.5000}, to: []float64{0.50, 0.00, 1.00}},
		{from: []float64{300.0, 1.0000, 0.5000}, to: []float64{1.00, 0.00, 1.00}},
		{from: []float64{330.0, 1.0000, 0.5000}, to: []float64{1.00, 0.00, 0.50}},
		{from: []float64{360.0, 1.0000, 0.5000}, to: []float64{1.00, 0.00, 0.00}},
		{from: []float64{000.0, 1.0000, 0.2500}, to: []float64{0.50, 0.00, 0.00}},
		{from: []float64{030.0, 1.0000, 0.2500}, to: []float64{0.50, 0.25, 0.00}},
		{from: []float64{060.0, 1.0000, 0.2500}, to: []float64{0.50, 0.50, 0.00}},
		{from: []float64{090.0, 1.0000, 0.2500}, to: []float64{0.25, 0.50, 0.00}},
		{from: []float64{120.0, 1.0000, 0.2500}, to: []float64{0.00, 0.50, 0.00}},
		{from: []float64{150.0, 1.0000, 0.2500}, to: []float64{0.00, 0.50, 0.25}},
		{from: []float64{180.0, 1.0000, 0.2500}, to: []float64{0.00, 0.50, 0.50}},
		{from: []float64{210.0, 1.0000, 0.2500}, to: []float64{0.00, 0.25, 0.50}},
		{from: []float64{240.0, 1.0000, 0.2500}, to: []float64{0.00, 0.00, 0.50}},
		{from: []float64{270.0, 1.0000, 0.2500}, to: []float64{0.25, 0.00, 0.50}},
		{from: []float64{300.0, 1.0000, 0.2500}, to: []float64{0.50, 0.00, 0.50}},
		{from: []float64{330.0, 1.0000, 0.2500}, to: []float64{0.50, 0.00, 0.25}},
		{from: []float64{360.0, 1.0000, 0.2500}, to: []float64{0.50, 0.00, 0.00}},
		{from: []float64{000.0, 1.0000, 0.7500}, to: []float64{1.00, 0.50, 0.50}},
		{from: []float64{030.0, 1.0000, 0.7500}, to: []float64{1.00, 0.75, 0.50}},
		{from: []float64{060.0, 1.0000, 0.7500}, to: []float64{1.00, 1.00, 0.50}},
		{from: []float64{090.0, 1.0000, 0.7500}, to: []float64{0.75, 1.00, 0.50}},
		{from: []float64{120.0, 1.0000, 0.7500}, to: []float64{0.50, 1.00, 0.50}},
		{from: []float64{150.0, 1.0000, 0.7500}, to: []float64{0.50, 1.00, 0.75}},
		{from: []float64{180.0, 1.0000, 0.7500}, to: []float64{0.50, 1.00, 1.00}},
		{from: []float64{210.0, 1.0000, 0.7500}, to: []float64{0.50, 0.75, 1.00}},
		{from: []float64{240.0, 1.0000, 0.7500}, to: []float64{0.50, 0.50, 1.00}},
		{from: []float64{270.0, 1.0000, 0.7500}, to: []float64{0.75, 0.50, 1.00}},
		{from: []float64{300.0, 1.0000, 0.7500}, to: []float64{1.00, 0.50, 1.00}},
		{from: []float64{330.0, 1.0000, 0.7500}, to: []float64{1.00, 0.50, 0.75}},
		{from: []float64{360.0, 1.0000, 0.7500}, to: []float64{1.00, 0.50, 0.50}},
	}

	for n := 0; n < len(tests); n++ {
		r, g, b, err := gocolor.HSLtoRGB(tests[n].from[0], tests[n].from[1], tests[n].from[2])

		assert.NoError(t, err)
		assert.InDeltaf(t, tests[n].to[0], r, precision, "r is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[1], g, precision, "g is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[2], b, precision, "b is wrong for test #%v", n+1)
	}
}

func TestHSVtoRGB_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-10, 0, 0},
		{390, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.HSVtoRGB(tests[n][0], tests[n][1], tests[n][2])
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestHSVtoRGB(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{000.0, 0.0000, 0.0000}, to: []float64{0.000, 0.000, 0.000}},
		{from: []float64{000.0, 0.0000, 1.0000}, to: []float64{1.000, 1.000, 1.000}},
		{from: []float64{000.0, 1.0000, 0.5000}, to: []float64{0.500, 0.000, 0.000}},
		{from: []float64{030.0, 1.0000, 0.5000}, to: []float64{0.500, 0.250, 0.000}},
		{from: []float64{060.0, 1.0000, 0.5000}, to: []float64{0.500, 0.500, 0.000}},
		{from: []float64{090.0, 1.0000, 0.5000}, to: []float64{0.250, 0.500, 0.000}},
		{from: []float64{120.0, 1.0000, 0.5000}, to: []float64{0.000, 0.500, 0.000}},
		{from: []float64{150.0, 1.0000, 0.5000}, to: []float64{0.000, 0.500, 0.250}},
		{from: []float64{180.0, 1.0000, 0.5000}, to: []float64{0.000, 0.500, 0.500}},
		{from: []float64{210.0, 1.0000, 0.5000}, to: []float64{0.000, 0.250, 0.500}},
		{from: []float64{240.0, 1.0000, 0.5000}, to: []float64{0.000, 0.000, 0.500}},
		{from: []float64{270.0, 1.0000, 0.5000}, to: []float64{0.250, 0.000, 0.500}},
		{from: []float64{300.0, 1.0000, 0.5000}, to: []float64{0.500, 0.000, 0.500}},
		{from: []float64{330.0, 1.0000, 0.5000}, to: []float64{0.500, 0.000, 0.250}},
		{from: []float64{360.0, 1.0000, 0.5000}, to: []float64{0.500, 0.000, 0.000}},
		{from: []float64{000.0, 1.0000, 0.2500}, to: []float64{0.250, 0.000, 0.000}},
		{from: []float64{030.0, 1.0000, 0.2500}, to: []float64{0.250, 0.125, 0.000}},
		{from: []float64{060.0, 1.0000, 0.2500}, to: []float64{0.250, 0.250, 0.000}},
		{from: []float64{090.0, 1.0000, 0.2500}, to: []float64{0.125, 0.250, 0.000}},
		{from: []float64{120.0, 1.0000, 0.2500}, to: []float64{0.000, 0.250, 0.000}},
		{from: []float64{150.0, 1.0000, 0.2500}, to: []float64{0.000, 0.250, 0.125}},
		{from: []float64{180.0, 1.0000, 0.2500}, to: []float64{0.000, 0.250, 0.250}},
		{from: []float64{210.0, 1.0000, 0.2500}, to: []float64{0.000, 0.125, 0.250}},
		{from: []float64{240.0, 1.0000, 0.2500}, to: []float64{0.000, 0.000, 0.250}},
		{from: []float64{270.0, 1.0000, 0.2500}, to: []float64{0.125, 0.000, 0.250}},
		{from: []float64{300.0, 1.0000, 0.2500}, to: []float64{0.250, 0.000, 0.250}},
		{from: []float64{330.0, 1.0000, 0.2500}, to: []float64{0.250, 0.000, 0.125}},
		{from: []float64{360.0, 1.0000, 0.2500}, to: []float64{0.250, 0.000, 0.000}},
		{from: []float64{000.0, 1.0000, 0.7500}, to: []float64{0.750, 0.000, 0.000}},
		{from: []float64{030.0, 1.0000, 0.7500}, to: []float64{0.750, 0.375, 0.000}},
		{from: []float64{060.0, 1.0000, 0.7500}, to: []float64{0.750, 0.750, 0.000}},
		{from: []float64{090.0, 1.0000, 0.7500}, to: []float64{0.375, 0.750, 0.000}},
		{from: []float64{120.0, 1.0000, 0.7500}, to: []float64{0.000, 0.750, 0.000}},
		{from: []float64{150.0, 1.0000, 0.7500}, to: []float64{0.000, 0.750, 0.375}},
		{from: []float64{180.0, 1.0000, 0.7500}, to: []float64{0.000, 0.750, 0.750}},
		{from: []float64{210.0, 1.0000, 0.7500}, to: []float64{0.000, 0.375, 0.750}},
		{from: []float64{240.0, 1.0000, 0.7500}, to: []float64{0.000, 0.000, 0.750}},
		{from: []float64{270.0, 1.0000, 0.7500}, to: []float64{0.375, 0.000, 0.750}},
		{from: []float64{300.0, 1.0000, 0.7500}, to: []float64{0.750, 0.000, 0.750}},
		{from: []float64{330.0, 1.0000, 0.7500}, to: []float64{0.750, 0.000, 0.375}},
		{from: []float64{360.0, 1.0000, 0.7500}, to: []float64{0.750, 0.000, 0.000}},
	}

	for n := 0; n < len(tests); n++ {
		r, g, b, err := gocolor.HSVtoRGB(tests[n].from[0], tests[n].from[1], tests[n].from[2])

		assert.NoError(t, err)
		assert.InDeltaf(t, tests[n].to[0], r, precision, "r is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[1], g, precision, "g is wrong for test #%v", n+1)
		assert.InDeltaf(t, tests[n].to[2], b, precision, "b is wrong for test #%v", n+1)
	}
}

func TestYIQtoRGB_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-1, 0, 0},
		{2, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.YIQtoRGB(tests[n][0], tests[n][1], tests[n][2])
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestYIQtoRGB(t *testing.T) {}

func TestYUVtoRGB(t *testing.T) {
	t.Run("convert SDYUV to RGB", TestSDYUVtoRGB)
	t.Run("convert HDYUV to RGB", TestHDYUVtoRGB)
}

func TestSDYUVtoRGB_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-1, 0, 0},
		{2, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.SDYUVtoRGB(tests[n][0], tests[n][1], tests[n][2])
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestSDYUVtoRGB(t *testing.T) {}

func TestHDYUVtoRGB_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-1, 0, 0},
		{2, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.HDYUVtoRGB(tests[n][0], tests[n][1], tests[n][2])
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestHDYUVtoRGB(t *testing.T) {}

func TestCMYtoRGB_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-1, 0, 0},
		{2, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.CMYtoRGB(tests[n][0], tests[n][1], tests[n][2])
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestCMYtoRGB(t *testing.T) {}

func TestHEXtoRGB_InvalidParameters(t *testing.T) {
	tests := []string {
		"aaaaaaaa",
		"gggggg",
		"fg0000",
		"00fg00",
		"0000fg",
		"g00",
		"0g0",
		"00g",
		"ff",
		"f",
		"",
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.HEXtoRGB(tests[n])
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestHEXtoRGB(t *testing.T) {}

func TestXYZtoRGB_InvalidParameters(t *testing.T) {
	tests := [][]float64{
		{-1, 0, 0},
		{2, 0, 0},
		{0, -1, 0},
		{0, 2, 0},
		{0, 0, -1},
		{0, 0, 2},
	}

	for n := 0; n < len(tests); n++ {
		_, _, _, err := gocolor.XYZtoRGB(tests[n][0], tests[n][1], tests[n][2], gocolor.AdobeRGB)
		assert.Errorf(t, err, "invalid parameter should return an error for test #%v", n+1)
	}
}

func TestXYZtoRGB(t *testing.T) {
	t.Run("convert XYZ to SRGB", TestXYZtoSRGB)
	t.Run("convert XYZ to BT2020", TestXYZtoBT2020)
	t.Run("convert XYZ to BT202012b", TestXYZtoBT202012b)
	t.Run("convert XYZ to AdobeRGB", TestXYZtoAdobeRGB)
}

func TestXYZtoSRGB(t *testing.T) {}

func TestXYZtoBT2020(t *testing.T) {}

func TestXYZtoBT202012b(t *testing.T) {}

func TestXYZtoAdobeRGB(t *testing.T) {}

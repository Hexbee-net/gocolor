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

func TestFromRGB(t *testing.T) {
	t.Run("convert RGB color to HSL", TestRGBtoHSL)
	t.Run("convert RGB color to HSV", TestRGBtoHSV)
	t.Run("convert RGB color to YIQ", TestRGBtoYIQ)
	t.Run("convert RGB color to YUV", TestRGBtoYUV)
	t.Run("convert RGB color to CMY", TestRGBtoCMY)
	t.Run("convert RGB color to HTML", TestRGBtoHTML)
	t.Run("convert RGB color to XYZ", TestRGBtoXYZ)
}

func TestRGBtoHSL(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{0, 0, 0}, to: []float64{0, 0, 0}},
		{from: []float64{1, 1, 1}, to: []float64{0, 0, 1}},
		{from: []float64{1, 0, 0}, to: []float64{0, 1, 0.5}},
		{from: []float64{0, 1, 0}, to: []float64{120, 1, 0.5}},
		{from: []float64{0, 0, 1}, to: []float64{240, 1, 0.5}},
		{from: []float64{1, 1, 0}, to: []float64{60, 1, 0.5}},
		{from: []float64{1, 0, 1}, to: []float64{300, 1, 0.5}},
		{from: []float64{0, 1, 1}, to: []float64{180, 1, 0.5}},
		{from: []float64{1, 0.5, 0}, to: []float64{30, 1, 0.5}},
		{from: []float64{1, 0.5, 0.5}, to: []float64{0, 1, 0.75}},
		{from: []float64{0.5, 1, 0}, to: []float64{90, 1, 0.5}},
		{from: []float64{0.5, 1, 0.5}, to: []float64{120, 1, 0.75}},
		{from: []float64{0.5, 0, 1}, to: []float64{270, 1, 0.5}},
		{from: []float64{0.5, 0.5, 1}, to: []float64{240, 1, 0.75}},
		{from: []float64{0.5, 0.5, 0.5}, to: []float64{0, 0, 0.5}},
		{from: []float64{0.5, 0.5, 0}, to: []float64{60, 1, 0.25}},
		{from: []float64{0.5, 0, 0.5}, to: []float64{300, 1, 0.25}},
		{from: []float64{0, 0.5, 0.5}, to: []float64{180, 1, 0.25}},
		{from: []float64{0.5, 0, 0}, to: []float64{0, 1, 0.25}},
		{from: []float64{0, 0.5, 0}, to: []float64{120, 1, 0.25}},
		{from: []float64{0, 0, 0.5}, to: []float64{240, 1, 0.25}},
		{from: []float64{1, 0.5, 0.25}, to: []float64{20, 1, 0.625}},
		{from: []float64{0.5, 1, 0.25}, to: []float64{100, 1, 0.625}},
		{from: []float64{0.5, 0.25, 1}, to: []float64{260, 1, 0.625}},
		{from: []float64{0.25, 1, 0.5}, to: []float64{140, 1, 0.625}},
		{from: []float64{0.25, 0.5, 1}, to: []float64{220, 1, 0.625}},
		{from: []float64{0.75, 0, 0}, to: []float64{0, 1, 0.375}},
		{from: []float64{0, 0.75, 0}, to: []float64{120, 1, 0.375}},
		{from: []float64{0, 0, 0.75}, to: []float64{240, 1, 0.375}},
		{from: []float64{0.75, 0.75, 0}, to: []float64{60, 1, 0.375}},
		{from: []float64{0.75, 0, 0.75}, to: []float64{300, 1, 0.375}},
		{from: []float64{0, 0.75, 0.75}, to: []float64{180, 1, 0.375}},
		{from: []float64{0.75, 0.5, 0}, to: []float64{40, 1, 0.375}},
		{from: []float64{0.75, 0, 0.5}, to: []float64{320, 1, 0.375}},
		{from: []float64{0.75, 0.5, 0.5}, to: []float64{0, 1.0 / 3.0, 0.625}},
		{from: []float64{0.5, 0.75, 0}, to: []float64{80, 1, 0.375}},
		{from: []float64{0, 0.75, 0.5}, to: []float64{160, 1, 0.375}},
		{from: []float64{0.5, 0.75, 0.5}, to: []float64{120, 1.0 / 3.0, 0.625}},
		{from: []float64{0.5, 0, 0.75}, to: []float64{280, 1, 0.375}},
		{from: []float64{0, 0.5, 0.75}, to: []float64{200, 1, 0.375}},
		{from: []float64{0.5, 0.5, 0.75}, to: []float64{240, 1.0 / 3.0, 0.625}},
	}

	for i := 0; i < len(tests); i++ {
		h, s, l := gocolor.RGBtoHSL(tests[i].from[0], tests[i].from[1], tests[i].from[2])

		assert.InDeltaf(t, tests[i].to[0], h, precision, "h is wrong for test #%v", i+1)
		assert.InDeltaf(t, tests[i].to[1], s, precision, "s is wrong for test #%v", i+1)
		assert.InDeltaf(t, tests[i].to[2], l, precision, "l is wrong for test #%v", i+1)
	}
}

func TestRGBtoHSV(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{0, 0, 0}, to: []float64{0, 0, 0}},
		{from: []float64{1, 1, 1}, to: []float64{0, 0, 1}},
		{from: []float64{1, 0, 0}, to: []float64{0, 1, 1}},
		{from: []float64{0, 1, 0}, to: []float64{120, 1, 1}},
		{from: []float64{0, 0, 1}, to: []float64{240, 1, 1}},
		{from: []float64{1, 1, 0}, to: []float64{60, 1, 1}},
		{from: []float64{1, 0, 1}, to: []float64{300, 1, 1}},
		{from: []float64{0, 1, 1}, to: []float64{180, 1, 1}},
		{from: []float64{1, 0.5, 0}, to: []float64{30, 1, 1}},
		{from: []float64{1, 0.5, 0.5}, to: []float64{0, 0.5, 1}},
		{from: []float64{0.5, 1, 0}, to: []float64{90, 1, 1}},
		{from: []float64{0.5, 1, 0.5}, to: []float64{120, 0.5, 1}},
		{from: []float64{0.5, 0, 1}, to: []float64{270, 1, 1}},
		{from: []float64{0.5, 0.5, 1}, to: []float64{240, 0.5, 1}},
		{from: []float64{0.5, 0.5, 0.5}, to: []float64{0, 0, 0.5}},
		{from: []float64{0.5, 0.5, 0}, to: []float64{60, 1, 0.5}},
		{from: []float64{0.5, 0, 0.5}, to: []float64{300, 1, 0.5}},
		{from: []float64{0, 0.5, 0.5}, to: []float64{180, 1, 0.5}},
		{from: []float64{0.5, 0, 0}, to: []float64{0, 1, 0.5}},
		{from: []float64{0, 0.5, 0}, to: []float64{120, 1, 0.5}},
		{from: []float64{0, 0, 0.5}, to: []float64{240, 1, 0.5}},
		{from: []float64{1, 0.5, 0.25}, to: []float64{20, 0.75, 1}},
		{from: []float64{0.5, 1, 0.25}, to: []float64{100, 0.75, 1}},
		{from: []float64{0.5, 0.25, 1}, to: []float64{260, 0.75, 1}},
		{from: []float64{0.25, 1, 0.5}, to: []float64{140, 0.75, 1}},
		{from: []float64{0.25, 0.5, 1}, to: []float64{220, 0.75, 1}},
		{from: []float64{0.75, 0, 0}, to: []float64{0, 1, 0.75}},
		{from: []float64{0, 0.75, 0}, to: []float64{120, 1, 0.75}},
		{from: []float64{0, 0, 0.75}, to: []float64{240, 1, 0.75}},
		{from: []float64{0.75, 0.75, 0}, to: []float64{60, 1, 0.75}},
		{from: []float64{0.75, 0, 0.75}, to: []float64{300, 1, 0.75}},
		{from: []float64{0, 0.75, 0.75}, to: []float64{180, 1, 0.75}},
		{from: []float64{0.75, 0.5, 0}, to: []float64{40, 1, 0.75}},
		{from: []float64{0.75, 0, 0.5}, to: []float64{320, 1, 0.75}},
		{from: []float64{0.75, 0.5, 0.5}, to: []float64{0, 1.0 / 3.0, 0.75}},
		{from: []float64{0.5, 0.75, 0}, to: []float64{80, 1, 0.75}},
		{from: []float64{0, 0.75, 0.5}, to: []float64{160, 1, 0.75}},
		{from: []float64{0.5, 0.75, 0.5}, to: []float64{120, 1.0 / 3.0, 0.75}},
		{from: []float64{0.5, 0, 0.75}, to: []float64{280, 1, 0.75}},
		{from: []float64{0, 0.5, 0.75}, to: []float64{200, 1, 0.75}},
		{from: []float64{0.5, 0.5, 0.75}, to: []float64{240, 1.0 / 3.0, 0.75}},
	}

	for i := 0; i < len(tests); i++ {
		h, s, v := gocolor.RGBtoHSV(tests[i].from[0], tests[i].from[1], tests[i].from[2])

		assert.InDeltaf(t, tests[i].to[0], h, precision, "h is wrong for test #%v", i+1)
		assert.InDeltaf(t, tests[i].to[1], s, precision, "s is wrong for test #%v", i+1)
		assert.InDeltaf(t, tests[i].to[2], v, precision, "v is wrong for test #%v", i+1)
	}
}

func TestRGBtoYIQ(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{0, 0, 0}, to: []float64{0, 0, 0}},
		{from: []float64{1, 1, 1}, to: []float64{1, 0, -1e-8}},
		{from: []float64{1, 0, 0}, to: []float64{0.29895808, 0.59590296, 0.21133576}},
		{from: []float64{0, 1, 0}, to: []float64{0.58660979, -0.27405705, -0.52263517}},
		{from: []float64{0, 0, 1}, to: []float64{0.11443213, -0.32184591, 0.31129940}},
		{from: []float64{1, 1, 0}, to: []float64{0.88556787, 0.32184591, -0.31129941}},
		{from: []float64{1, 0, 1}, to: []float64{0.41339021, 0.27405705, 0.52263516}},
		{from: []float64{0, 1, 1}, to: []float64{0.70104192, -0.59590296, -0.21133577}},
		{from: []float64{1, 0.5, 0}, to: []float64{0.59226298, 0.45887444, -0.04998183}},
		{from: []float64{1, 0.5, 0.5}, to: []float64{0.64947904, 0.29795148, 0.10566788}},
		{from: []float64{0.5, 1, 0}, to: []float64{0.73608883, 0.02389443, -0.41696729}},
		{from: []float64{0.5, 1, 0.5}, to: []float64{0.79330490, -0.13702853, -0.26131759}},
		{from: []float64{0.5, 0, 1}, to: []float64{0.26391117, -0.02389443, 0.41696728}},
		{from: []float64{0.5, 0.5, 1}, to: []float64{0.55721607, -0.16092296, 0.15564970}},
		{from: []float64{0.5, 0.5, 0.5}, to: []float64{0.50000000, 0.00000000, 0.00000000}},
		{from: []float64{0.5, 0.5, 0}, to: []float64{0.44278394, 0.16092296, -0.15564971}},
		{from: []float64{0.5, 0, 0.5}, to: []float64{0.20669511, 0.13702853, 0.26131758}},
		{from: []float64{0, 0.5, 0.5}, to: []float64{0.35052096, -0.29795148, -0.10566789}},
		{from: []float64{0.5, 0, 0}, to: []float64{0.14947904, 0.29795148, 0.10566788}},
		{from: []float64{0, 0.5, 0}, to: []float64{0.29330490, -0.13702853, -0.26131759}},
		{from: []float64{0, 0, 0.5}, to: []float64{0.05721607, -0.16092296, 0.15564970}},
		{from: []float64{1, 0.5, 0.25}, to: []float64{0.62087101, 0.37841296, 0.02784303}},
		{from: []float64{0.5, 1, 0.25}, to: []float64{0.76469686, -0.05656705, -0.33914244}},
		{from: []float64{0.5, 0.25, 1}, to: []float64{0.41056362, -0.09240869, 0.28630849}},
		{from: []float64{0.25, 1, 0.5}, to: []float64{0.71856538, -0.28600427, -0.31415153}},
		{from: []float64{0.25, 0.5, 1}, to: []float64{0.48247655, -0.30989870, 0.10281576}},
		{from: []float64{0.75, 0, 0}, to: []float64{0.22421856, 0.44692722, 0.15850182}},
		{from: []float64{0, 0.75, 0}, to: []float64{0.43995734, -0.20554279, -0.39197638}},
		{from: []float64{0, 0, 0.75}, to: []float64{0.08582410, -0.24138443, 0.23347455}},
		{from: []float64{0.75, 0.75, 0}, to: []float64{0.66417590, 0.24138443, -0.23347456}},
		{from: []float64{0.75, 0, 0.75}, to: []float64{0.31004266, 0.20554279, 0.39197637}},
		{from: []float64{0, 0.75, 0.75}, to: []float64{0.52578144, -0.44692722, -0.15850183}},
		{from: []float64{0.75, 0.5, 0}, to: []float64{0.51752346, 0.30989870, -0.10281577}},
		{from: []float64{0.75, 0, 0.5}, to: []float64{0.28143462, 0.28600427, 0.31415152}},
		{from: []float64{0.75, 0.5, 0.5}, to: []float64{0.57473952, 0.14897574, 0.05283394}},
		{from: []float64{0.5, 0.75, 0}, to: []float64{0.58943638, 0.09240869, -0.28630850}},
		{from: []float64{0, 0.75, 0.5}, to: []float64{0.49717341, -0.36646574, -0.23632668}},
		{from: []float64{0.5, 0.75, 0.5}, to: []float64{0.64665245, -0.06851426, -0.13065880}},
		{from: []float64{0.5, 0, 0.75}, to: []float64{0.23530314, 0.05656705, 0.33914243}},
		{from: []float64{0, 0.5, 0.75}, to: []float64{0.37912899, -0.37841296, -0.02784304}},
		{from: []float64{0.5, 0.5, 0.75}, to: []float64{0.52860803, -0.08046148, 0.07782485}},
	}

	for n := 0; n < len(tests); n++ {
		y, i, q := gocolor.RGBtoYIQ(tests[n].from[0], tests[n].from[1], tests[n].from[2])

		assert.InDeltaf(t, tests[n].to[0], y, precision, "y is wrong for test #%v", i+1)
		assert.InDeltaf(t, tests[n].to[1], i, precision, "i is wrong for test #%v", i+1)
		assert.InDeltaf(t, tests[n].to[2], q, precision, "q is wrong for test #%v", i+1)
	}
}

func TestRGBtoYUV(t *testing.T) {
}

func TestRGBtoCMY(t *testing.T) {
	tests := []ConversionTest{
		{from: []float64{0, 0, 0}, to: []float64{1, 1, 1}},
		{from: []float64{1, 1, 1}, to: []float64{0, 0, 0}},
		{from: []float64{1, 0, 0}, to: []float64{0, 1, 1}},
		{from: []float64{0, 1, 0}, to: []float64{1, 0, 1}},
		{from: []float64{0, 0, 1}, to: []float64{1, 1, 0}},
		{from: []float64{1, 1, 0}, to: []float64{0, 0, 1}},
		{from: []float64{1, 0, 1}, to: []float64{0, 1, 0}},
		{from: []float64{0, 1, 1}, to: []float64{1, 0, 0}},
		{from: []float64{1, 0.5, 0}, to: []float64{0, 0.5, 1}},
		{from: []float64{1, 0.5, 0.5}, to: []float64{0, 0.5, 0.5}},
		{from: []float64{0.5, 1, 0}, to: []float64{0.5, 0, 1}},
		{from: []float64{0.5, 1, 0.5}, to: []float64{0.5, 0, 0.5}},
		{from: []float64{0.5, 0, 1}, to: []float64{0.5, 1, 0}},
		{from: []float64{0.5, 0.5, 1}, to: []float64{0.5, 0.5, 0}},
		{from: []float64{0.5, 0.5, 0.5}, to: []float64{0.5, 0.5, 0.5}},
		{from: []float64{0.5, 0.5, 0}, to: []float64{0.5, 0.5, 1}},
		{from: []float64{0.5, 0, 0.5}, to: []float64{0.5, 1, 0.5}},
		{from: []float64{0, 0.5, 0.5}, to: []float64{1, 0.5, 0.5}},
		{from: []float64{0.5, 0, 0}, to: []float64{0.5, 1, 1}},
		{from: []float64{0, 0.5, 0}, to: []float64{1, 0.5, 1}},
		{from: []float64{0, 0, 0.5}, to: []float64{1, 1, 0.5}},
		{from: []float64{1, 0.5, 0.25}, to: []float64{0, 0.5, 0.75}},
		{from: []float64{0.5, 1, 0.25}, to: []float64{0.5, 0, 0.75}},
		{from: []float64{0.5, 0.25, 1}, to: []float64{0.5, 0.75, 0}},
		{from: []float64{0.25, 1, 0.5}, to: []float64{0.75, 0, 0.5}},
		{from: []float64{0.25, 0.5, 1}, to: []float64{0.75, 0.5, 0}},
		{from: []float64{0.75, 0, 0}, to: []float64{0.25, 1, 1}},
		{from: []float64{0, 0.75, 0}, to: []float64{1, 0.25, 1}},
		{from: []float64{0, 0, 0.75}, to: []float64{1, 1, 0.25}},
		{from: []float64{0.75, 0.75, 0}, to: []float64{0.25, 0.25, 1}},
		{from: []float64{0.75, 0, 0.75}, to: []float64{0.25, 1, 0.25}},
		{from: []float64{0, 0.75, 0.75}, to: []float64{1, 0.25, 0.25}},
		{from: []float64{0.75, 0.5, 0}, to: []float64{0.25, 0.5, 1}},
		{from: []float64{0.75, 0, 0.5}, to: []float64{0.25, 1, 0.5}},
		{from: []float64{0.75, 0.5, 0.5}, to: []float64{0.25, 0.5, 0.5}},
		{from: []float64{0.5, 0.75, 0}, to: []float64{0.5, 0.25, 1}},
		{from: []float64{0, 0.75, 0.5}, to: []float64{1, 0.25, 0.5}},
		{from: []float64{0.5, 0.75, 0.5}, to: []float64{0.5, 0.25, 0.5}},
		{from: []float64{0.5, 0, 0.75}, to: []float64{0.5, 1, 0.25}},
		{from: []float64{0, 0.5, 0.75}, to: []float64{1, 0.5, 0.25}},
		{from: []float64{0.5, 0.5, 0.75}, to: []float64{0.5, 0.5, 0.25}},
	}

	for i := 0; i < len(tests); i++ {
		c, m, y := gocolor.RGBtoCMY(tests[i].from[0], tests[i].from[1], tests[i].from[2])

		assert.InDeltaf(t, tests[i].to[0], c, precision, "c is wrong for test #%v", i+1)
		assert.InDeltaf(t, tests[i].to[1], m, precision, "m is wrong for test #%v", i+1)
		assert.InDeltaf(t, tests[i].to[2], y, precision, "y is wrong for test #%v", i+1)
	}
}

func TestRGBtoHTML(t *testing.T) {
}

func TestRGBtoXYZ(t *testing.T) {
}

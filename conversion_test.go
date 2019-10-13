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

const precision = 1e-10

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
}

func TestRGBtoYUV(t *testing.T) {
}

func TestRGBtoCMY(t *testing.T) {
}

func TestRGBtoHTML(t *testing.T) {
}

func TestRGBtoXYZ(t *testing.T) {
}

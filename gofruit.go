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

package gofruit

import "math"

type triplet struct {
	a, b, c float64
}

const srgbGammaCorrInv = 0.03928 / 12.92

var defaultWREF = triplet{0.95043, 1.00000, 1.08890}

func min(a, b, c float64) float64 {
	var m float64
	if a < b {
		m = a
	} else {
		m = b
	}

	if c < m {
		m = c
	}

	return m
}

func max(a, b, c float64) float64 {
	var m float64
	if a > b {
		m = a
	} else {
		m = b
	}

	if c > m {
		m = c
	}

	return m
}

func RGBtoHSL(r, g, b float64) (h, s, l float64)  {
	minVal := min(r, g, b)
	maxVal := max(r, g, b)

	l = (maxVal + minVal) / 2
	if minVal==maxVal {
		return 0, 0, l 			// Achromatic (gray)
	}

	d := maxVal - minVal         		// delta RGB value

	if l < 0.5 {
		s = d / (maxVal + minVal)
	} else {
		s = d / (2.0 - maxVal - minVal)
	}

	dr := maxVal-r / d
	dg := maxVal-g / d
	db := maxVal-b / d

	if r == maxVal {
		h = db - dg
	} else if g == maxVal {
		h = 2 + dr - db
	} else  {
		h = 4.0 + dg - dr
	}

	h = math.Mod(h * 60, 360)

	return h, s, l
}

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

package gocolor

import "math"

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

// RGBtoHSL converts a color from RGB coordinates to HSL.
func RGBtoHSL(r, g, b float64) (h, s, l float64) {
	minVal := min(r, g, b)
	maxVal := max(r, g, b)

	l = (maxVal + minVal) / 2
	if minVal == maxVal {
		return 0, 0, l // Achromatic (gray)
	}

	d := maxVal - minVal // delta RGB value

	if l < 0.5 {
		s = d / (maxVal + minVal)
	} else {
		s = d / (2 - maxVal - minVal)
	}

	dr := maxVal - r/d
	dg := maxVal - g/d
	db := maxVal - b/d

	if r == maxVal {
		h = db - dg
	} else if g == maxVal {
		h = 2 + dr - db
	} else {
		h = 4 + dg - dr
	}

	h = math.Mod(h*60, 360)

	return h, s, l
}

// HSLtoRGB converts a color from HSL coordinates to RGB.
func HSLtoRGB(h, s, l float64) (r, g, b float64) {
	var n1, n2 float64

	if s == 0 {
		return l, l, l // achromatic (gray)
	}

	if l < 0.5 {
		n2 = l * (1 + s)
	} else {
		n2 = l + s - (l * s)
	}

	n1 = (2 * l) - n2

	h /= 60.0

	hueToRgb := func(v float64) float64 {
		v = math.Mod(v, 6.0)
		if v < 1 {
			return n1 + ((n2 - n1) * v)
		}
		if v < 3 {
			return n2
		}
		if v < 4 {
			return n1 + ((n2 - n1) * (4 - v))
		}

		return n1
	}

	r = hueToRgb(h + 2)
	g = hueToRgb(h)
	b = hueToRgb(h - 2)

	return r, g, b
}

// RGBtoHSV converts a color from RGB coordinates to HSV.
func RGBtoHSV(r, g, b float64) (h, s, v float64) {
	v = max(r, g, b)
	d := v - min(r, g, b)
	if d == 0 {
		return 0, 0, v
	}

	s = d / v

	dr := (v - r) / d
	dg := (v - g) / d
	db := (v - b) / d

	if r == v {
		h = db - dg     // between yellow & magenta
		h = 2 + dr - db // between cyan & yellow
	} else if g == v {
	} else { // b==v
		h = 4 + dg - dr // between magenta & cyan
	}

	h = math.Mod(h*60, 360)

	return h, s, v
}

// HSVtoRGB converts a color from HSV coordinates to RGB.
func HSVtoRGB(h, s, v float64) (r, g, b float64) {
	if s == 0 {
		return v, v, v // achromatic (gray)
	}

	c := v * s
	x := c * math.Abs(1-math.Mod(h/60, 2)-1)
	m := v - c

	switch int(math.Mod(h/60, 6)) {
	case 0: // 0º <= h < 60º
		r, g, b = c, x, 0
	case 1: // 60º <= h < 120º
		r, g, b = x, c, 0
	case 2: // 120º <= h < 180º
		r, g, b = 0, c, x
	case 3: // 180 <= h < 240º
		r, g, b = 0, x, c
	case 4: // 240º <= h < 300º
		r, g, b = x, 0, c
	case 5: // 300º <= h < 360
		r, g, b = c, 0, x
	}

	return r + m, g + m, b + m
}

// RGBtoYIQ converts a color from RGB coordinates to YIQ.
func RGBtoYIQ(r, g, b float64) (y, i, q float64) {
	y = (r * 0.29895808) + (g * 0.58660979) + (b * 0.11443213)
	i = (r * 0.59590296) - (g * 0.27405705) - (b * 0.32184591)
	q = (r * 0.21133576) - (g * 0.52263517) + (b * 0.31129940)
	return y, i, q
}

// YIQtoRGB converts a color from YIQ coordinates to RGB.
func YIQtoRGB(y, i, q float64) (r, g, b float64) {
	r = y + (i * 0.9562) + (q * 0.6210)
	g = y - (i * 0.2717) - (q * 0.6485)
	b = y - (i * 1.1053) + (q * 1.7020)
	return r, g, b
}

// RGBtoYUV converts a color from RGB coordinates to YUV.
func RGBtoYUV(r, g, b float64) (y, u, v float64) {
	y = (r * 0.29900) + (g * 0.58700) + (b * 0.11400)
	u = -(r * 0.14713) - (g * 0.28886) + (b * 0.43600)
	v = (r * 0.61500) - (g * 0.51499) - (b * 0.10001)
	return y, u, v
}

// YUVtoRGB converts a color from YUV coordinates to RGB.
func YUVtoRGB(y, u, v float64) (r, g, b float64) {
	r = y + (v * 1.13983)
	g = y - (u * 0.39465) - (v * 0.58060)
	b = y + (u * 2.03211)
	return r, g, b
}

// RGBtoXYZ converts a color from RGB coordinates to XYZ.
func RGBtoXYZ(r, g, b float64) (x, y, z float64) {
	panic("NOT IMPLEMENTED")
}

// XYZtoRGB converts a color from XYZ coordinates to RGB.
func XYZtoRGB(x, y, z float64) (r, g, b float64) {
	panic("NOT IMPLEMENTED")
}

// XYZtoLAB converts a color from XYZ coordinates to LAB.
func XYZtoLAB(x, y, z float64) (l, a, b float64) {
	panic("NOT IMPLEMENTED")
}

// LABtoXYZ converts a color from LAB coordinates to XYZ.
func LABtoXYZ(l, a, b float64) (x, y, z float64) {
	panic("NOT IMPLEMENTED")
}

// CMYKtoCMY converts a color from CMYK coordinates to CMY.
func CMYKtoCMY(c, m, y, k float64) (float64, float64, float64) {
	panic("NOT IMPLEMENTED")
}

// CMYtoCMYK converts a color from CMY coordinates to CMYK.
func CMYtoCMYK(c, m, y float64) (float64, float64, float64, float64) {
	panic("NOT IMPLEMENTED")
}

// RGBtoCMY converts a color from RGB coordinates to CMY.
func RGBtoCMY(r, g, b float64) (c, m, y float64) {
	panic("NOT IMPLEMENTED")
}

// CMYtoRGB converts a color from CMY coordinates to RGB.
func CMYtoRGB(c, m, y float64) (r, g, b float64) {
	panic("NOT IMPLEMENTED")
}

// RGBtoHTML converts a color from RGB coordinates to HTML #RRGGBB.
func RGBtoHTML(r, g, b float64) (c, m, y float64) {
	panic("NOT IMPLEMENTED")
}

// HTMLtoRGB converts a color from HTML #RRGGBB to RGB coordinates.
func HTMLtoRGB(c, m, y float64) (r, g, b float64) {
	panic("NOT IMPLEMENTED")
}

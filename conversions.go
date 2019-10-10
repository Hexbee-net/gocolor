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

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Hexbee-net/gocolor/colorspace"
	"github.com/Hexbee-net/gocolor/named"
)

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

////////////////////////////////////////

// RGBtoHSL converts a color from base RGB coordinates to HSL.
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

// RGBtoHSV converts a color from base RGB coordinates to HSV.
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

// RGBtoYIQ converts a color from base RGB coordinates to YIQ.
func RGBtoYIQ(r, g, b float64) (y, i, q float64) {
	y = (r * 0.29895808) + (g * 0.58660979) + (b * 0.11443213)
	i = (r * 0.59590296) - (g * 0.27405705) - (b * 0.32184591)
	q = (r * 0.21133576) - (g * 0.52263517) + (b * 0.31129940)
	return y, i, q
}

// RGBtoYUV converts a color from base RGB coordinates to YUV.
func RGBtoYUV(r, g, b float64) (y, u, v float64) {
	y = (r * 0.29900) + (g * 0.58700) + (b * 0.11400)
	u = -(r * 0.14713) - (g * 0.28886) + (b * 0.43600)
	v = (r * 0.61500) - (g * 0.51499) - (b * 0.10001)
	return y, u, v
}

// RGBtoCMY converts a color from base RGB coordinates to CMY.
func RGBtoCMY(r, g, b float64) (float64, float64, float64) {
	return 1 - r, 1 - g, 1 - b
}

// RGBtoHTML converts a color from base RGB coordinates to HTML #RRGGBB.
func RGBtoHTML(r, g, b float64) string {
	ri := int(math.Min(math.Round(r*255), 255))
	gi := int(math.Min(math.Round(g*255), 255))
	bi := int(math.Min(math.Round(b*255), 255))
	return fmt.Sprintf("#%02X%02X%02X", ri, gi, bi)
}

// RGBtoXYZ converts a color from RGB coordinates to XYZ.
func RGBtoXYZ(r, g, b float64, space string) (x, y, z float64) {
	switch space {
	case colorspace.SRGB:
		linearize := func(v float64) float64 {
			if v <= 0.04045 {
				return v / 12.92
			}
			return math.Pow((v+0.055)/1.055, 2.4)
		}
		r = linearize(r)
		g = linearize(g)
		b = linearize(b)

	case colorspace.BT2020:
		linearize := func(v float64) float64 {
			if v <= 0.08124794403514049 {
				return v / 4.5
			}
			return math.Pow((v+0.099)/1.099, 1/0.45)
		}
		r = linearize(r)
		g = linearize(g)
		b = linearize(b)

	case colorspace.BT202012b:
		linearize := func(v float64) float64 {
			if v <= 0.081697877417347 {
				return v / 4.5
			}
			return math.Pow((v+0.0993)/1.0993, 1/0.45)
		}
		r = linearize(r)
		g = linearize(g)
		b = linearize(b)

	default:
		gamma, ok := colorspace.Gamma[space]
		if !ok {
			panic(fmt.Sprintf("unrecognized RGB color space: %v", space))
		}
		r = math.Pow(r, gamma)
		g = math.Pow(g, gamma)
		b = math.Pow(b, gamma)
	}

	m, ok := conversionRgbXyz[space]
	if !ok {
		panic(fmt.Sprintf("unrecognized RGB color space: %v", space))
	}

	v := m.vdot(vector{r, g, b})
	return math.Max(v.v0, 0), math.Max(v.v1, 0), math.Max(v.v2, 0)
}

////////////////////////////////////////

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

// YIQtoRGB converts a color from YIQ coordinates to RGB.
func YIQtoRGB(y, i, q float64) (r, g, b float64) {
	r = y + (i * 0.9562) + (q * 0.6210)
	g = y - (i * 0.2717) - (q * 0.6485)
	b = y - (i * 1.1053) + (q * 1.7020)
	return r, g, b
}

// YUVtoRGB converts a color from YUV coordinates to RGB.
func YUVtoRGB(y, u, v float64) (r, g, b float64) {
	r = y + (v * 1.13983)
	g = y - (u * 0.39465) - (v * 0.58060)
	b = y + (u * 2.03211)
	return r, g, b
}

// CMYtoRGB converts a color from CMY coordinates to RGB.
func CMYtoRGB(c, m, y float64) (float64, float64, float64) {
	return 1 - c, 1 - m, 1 - y
}

// HTMLtoRGB converts a color from HTML #RRGGBB to RGB coordinates.
func HTMLtoRGB(html string) (r, g, b float64) {
	html = strings.TrimSpace(html)
	if html[0] == '#' {
		html = html[1:]
	} else {
		if val, ok := named.NamedColors[strings.ToLower(html)]; ok {
			html = val[1:]
		}
	}

	switch len(html) {
	// Long html code
	case 6:
		ri, err := strconv.ParseUint(html[0:2], 16, 64)
		if err != nil {
			panic(err)
		}
		r = float64(ri) / 255

		gi, err := strconv.ParseUint(html[2:4], 16, 64)
		if err != nil {
			panic(err)
		}
		g = float64(gi) / 255

		bi, err := strconv.ParseUint(html[4:6], 16, 64)
		if err != nil {
			panic(err)
		}
		b = float64(bi) / 255

	// Short html code
	case 3:
		ri, err := strconv.ParseUint(html[0:1], 16, 64)
		if err != nil {
			panic(err)
		}
		r = float64(ri*16+ri) / 255

		gi, err := strconv.ParseUint(html[1:2], 16, 64)
		if err != nil {
			panic(err)
		}
		g = float64(gi*16+gi) / 255

		bi, err := strconv.ParseUint(html[2:3], 16, 64)
		if err != nil {
			panic(err)
		}
		b = float64(bi*16+bi) / 255

	default:
		panic(fmt.Sprintf("input '%s' is not in #RRGGBB format", html))
	}

	return r, g, b
}

// XYZtoRGB converts a color from XYZ coordinates to RGB.
func XYZtoRGB(x, y, z float64, space string) (r, g, b float64) {
	m, ok := conversionXyzRgb[space]
	if !ok {
		panic(fmt.Sprintf("unrecognized RGB color space: %v", space))
	}

	v := m.vdot(vector{x, y, z})
	r, g, b = v.v0, v.v1, v.v2

	switch space {
	case colorspace.SRGB:
		delinearize := func(v float64) float64 {
			if v <= 0.0031308 {
				return v * 12.92
			}
			return 1.055*math.Pow(v, 1/2.4) - 0.055
		}
		r = delinearize(r)
		g = delinearize(g)
		b = delinearize(b)

	case colorspace.BT2020:
		delinearize := func(v float64) float64 {
			if v < 0.018 {
				return v * 4.5
			}
			return 1.099*math.Pow(v, 0.45) - 0.099
		}
		r = delinearize(r)
		g = delinearize(g)
		b = delinearize(b)

	case colorspace.BT202012b:
		delinearize := func(v float64) float64 {
			if v < 0.0181 {
				return v * 4.5
			}
			return 1.0993*math.Pow(v, 0.45) - 0.0993
		}
		r = delinearize(r)
		g = delinearize(g)
		b = delinearize(b)

	default:
		gamma, ok := colorspace.Gamma[space]
		if !ok {
			panic(fmt.Sprintf("unrecognized RGB color space: %v", space))
		}
		r = math.Pow(r, 1/gamma)
		g = math.Pow(g, 1/gamma)
		b = math.Pow(b, 1/gamma)
	}

	return r, g, b
}

////////////////////////////////////////

// CMYKtoCMY converts a color from CMYK coordinates to CMY.
func CMYKtoCMY(c, m, y, k float64) (float64, float64, float64) {
	mk := 1 - k
	return (c * mk) + k, (m * mk) + k, (y * mk) + k
}

// CMYtoCMYK converts a color from CMY coordinates to CMYK.
func CMYtoCMYK(c, m, y float64) (float64, float64, float64, float64) {
	k := min(c, m, y)
	if k == 1.0 {
		return 0.0, 0.0, 0.0, 1.0
	}

	mk := 1 - k
	return (c - k) / mk, (m - k) / mk, (y - k) / mk, k
}

////////////////////////////////////////

// XYZtoLAB converts a color from XYZ coordinates to LAB.
func XYZtoLAB(x, y, z float64) (l, a, b float64) {
	panic("NOT IMPLEMENTED")
}

// XYZtoXYY converts a color from XYZ coordinates to xyY.
func XYZtoXYY() {
	panic("NOT IMPLEMENTED")
}

// XYZtoLUV converts a color from XYZ coordinates to Luv.
func XYZtoLUV() {
	panic("NOT IMPLEMENTED")
}

// XYYtoXYZ converts a color from xyZ coordinates to XYZ.
func XYYtoXYZ() {
	panic("NOT IMPLEMENTED")
}

// LUVtoXYZ converts a color from Luv coordinates to XYZ.
func LUVtoXYZ() {
	panic("NOT IMPLEMENTED")
}

// XYZtoIPT converts a color from XYZ coordinates to IPT.
func XYZtoIPT(x, y, z float64) (i, p, t float64) {
	panic("NOT IMPLEMENTED")
}

////////////////////////////////////////

// LABtoXYZ converts a color from LAB coordinates to XYZ.
func LABtoXYZ(l, a, b float64) (x, y, z float64) {
	panic("NOT IMPLEMENTED")
}

// LABtoLCHAB converts a color from LAB coordinates to LCHab.
func LABtoLCHAB() {
	panic("NOT IMPLEMENTED")
}

// LCHABtoLAB converts a color from LCHab coordinates to LAB.
func LCHABtoLAB() {
	panic("NOT IMPLEMENTED")
}

// LUVtoLCHUV converts a color from Luv coordinates to LCHuv.
func LUVtoLCHUV() {
	panic("NOT IMPLEMENTED")
}

// LCHUVtoLUV converts a color from LCHuv coordinates to Luv.
func LCHUVtoLUV() {
	panic("NOT IMPLEMENTED")
}

////////////////////////////////////////

// IPTtoXYZ converts a color from IPT coordinates to XYZ.
func IPTtoXYZ(i, p, t float64) (x, y, z float64) {
	panic("NOT IMPLEMENTED")
}

// SpectralToXYZ converts spectral readings to XYZ coordinates.
func SpectralToXYZ(color []float64, observer Observer, refIlluminant []float64) (x, y, z float64) {
	var (
		stdObserverX = stdObs10X
		stdObserverY = stdObs10Y
		stdObserverZ = stdObs10Z
	)

	if observer == Observer2 {
		stdObserverX = stdObs2X
		stdObserverY = stdObs2Y
		stdObserverZ = stdObs2Z
	}

	l := len(color)
	if l != len(stdObserverX) || l != len(refIlluminant) {
		panic("mismatching spectral sampling length")
	}

	var (
		denom      float64 = 0
		xNumerator float64 = 0
		yNumerator float64 = 0
		zNumerator float64 = 0
	)
	for i := 0; i < l; i++ {
		denom += stdObserverY[i] * refIlluminant[i]

		sampleByRefIlluminant := color[i] * refIlluminant[i]
		xNumerator += sampleByRefIlluminant * stdObserverX[i]
		yNumerator += sampleByRefIlluminant * stdObserverY[i]
		zNumerator += sampleByRefIlluminant * stdObserverZ[i]
	}

	x = xNumerator / denom
	y = yNumerator / denom
	z = zNumerator / denom

	return x, y, z
}

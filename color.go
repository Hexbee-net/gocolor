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

// Observer is used to define a standard observer angle of vision.
type Observer int

// Reference illuminants
const (
	RefIlluminantA         = "a"
	RefIlluminantB         = "b"
	RefIlluminantC         = "c"
	RefIlluminantD50       = "d50"
	RefIlluminantD55       = "d55"
	RefIlluminantD65       = "d65"
	RefIlluminantD75       = "d75"
	RefIlluminantE         = "e"
	RefIlluminantF2        = "f2"
	RefIlluminantF7        = "f7"
	RefIlluminantF11       = "f11"
	RefIlluminantBlackBody = "blackbody"
)

// Standard observers
const (
	Observer2  = 2
	Observer10 = 10
)

// // Illuminant stores information on the color illuminant and observer angle.
// //
// // Illuminants and observer angles are used in all color spaces that use
// // reflective (instead of transmissive) light.
// type Illuminant struct {
// 	Observer   Observer
// 	Illuminant string
// }

// Wavelength is used to index the differrent wavelengths in a spectral color.
type Wavelength int

const (
	// Ultraviolet
	spec340nm Wavelength = iota
	spec350nm Wavelength = iota
	spec360nm Wavelength = iota
	spec370nm Wavelength = iota
	spec380nm Wavelength = iota

	// Violet
	spec390nm Wavelength = iota
	spec400nm Wavelength = iota
	spec410nm Wavelength = iota
	spec420nm Wavelength = iota
	spec430nm Wavelength = iota
	spec440nm Wavelength = iota
	spec450nm Wavelength = iota

	// Blue
	spec460nm Wavelength = iota
	spec470nm Wavelength = iota
	spec480nm Wavelength = iota
	spec490nm Wavelength = iota

	// Green
	spec500nm Wavelength = iota
	spec510nm Wavelength = iota
	spec520nm Wavelength = iota
	spec530nm Wavelength = iota
	spec540nm Wavelength = iota
	spec550nm Wavelength = iota
	spec560nm Wavelength = iota
	spec570nm Wavelength = iota

	// Yellow
	spec580nm Wavelength = iota
	spec590nm Wavelength = iota

	// Orange
	spec600nm Wavelength = iota
	spec610nm Wavelength = iota

	// Red
	spec620nm Wavelength = iota
	spec630nm Wavelength = iota
	spec640nm Wavelength = iota
	spec650nm Wavelength = iota
	spec660nm Wavelength = iota
	spec670nm Wavelength = iota
	spec680nm Wavelength = iota
	spec690nm Wavelength = iota
	spec700nm Wavelength = iota

	// Infrared
	spec710nm Wavelength = iota
	spec720nm Wavelength = iota
	spec730nm Wavelength = iota
	spec740nm Wavelength = iota
	spec750nm Wavelength = iota
	spec760nm Wavelength = iota
	spec770nm Wavelength = iota
	spec780nm Wavelength = iota
	spec790nm Wavelength = iota
	spec800nm Wavelength = iota
	spec810nm Wavelength = iota
	spec820nm Wavelength = iota
	spec830nm Wavelength = iota
)

// SpectralColor represents a spectral power distribution, as read by
// a spectrophotometer.
// The library assumes wavelength intervals of 10nm, starting at 340nm and ending at 830nm.
//
// Spectral colors are the lowest level, most "raw" measurement of color.
// You may convert spectral colors to any other color space, but you can't
// convert any other color space back to spectral.
//
// See `Spectral power distribution http://en.wikipedia.org/wiki/Spectral_power_distribution
// on Wikipedia for some higher level details on how these work.
type SpectralColor []float64

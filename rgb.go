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

const (
	AdobeRGB      = "Adobe RGB"
	AppleRGB      = "AppleRGB"
	BestRGB       = "Best RGB"
	BetaRGB       = "Beta RGB"
	BruceRGB      = "Bruce RGB"
	BT2020        = "ITU-R BT.2020"
	BT202012b     = "ITU-R BT.2020 12 bits"
	CieRGB        = "CIE RGB"
	ColorMatchRGB = "ColorMatch RGB"
	DonRGB4       = "Don RGB 4"
	EciRGB        = "ECI RGB"
	EktaSpacePS5  = "Ekta Space PS5"
	NtscRGB       = "NTSC RGB"
	PalSecamRGB   = "PAL/SECAM RGB"
	ProPhotoRGB   = "ProPhoto RGB"
	SmptecRGB     = "SMPTE-C RGB"
	SRGB          = "sRgb"
	WideGamutRGB  = "Wide Gamut RGB"
)

var RGBGamma = map[string]float64{
	AdobeRGB:      2.2,
	AppleRGB:      1.8,
	BestRGB:       2.2,
	BetaRGB:       2.2,
	BruceRGB:      2.2,
	CieRGB:        2.2,
	ColorMatchRGB: 1.8,
	DonRGB4:       2.2,
	EktaSpacePS5:  2.2,
	NtscRGB:       2.2,
	PalSecamRGB:   2.2,
	ProPhotoRGB:   1.8,
	SmptecRGB:     2.2,
	SRGB:          2.2,
	WideGamutRGB:  2.2,
}

var RGBIlluminants = map[string]string{
	AdobeRGB:      RefIlluminantD65,
	AppleRGB:      RefIlluminantD65,
	BestRGB:       RefIlluminantD50,
	BetaRGB:       RefIlluminantD50,
	BruceRGB:      RefIlluminantD65,
	BT2020:        RefIlluminantD65,
	CieRGB:        RefIlluminantE,
	ColorMatchRGB: RefIlluminantD50,
	DonRGB4:       RefIlluminantD50,
	EktaSpacePS5:  RefIlluminantD50,
	NtscRGB:       RefIlluminantD50,
	PalSecamRGB:   RefIlluminantD65,
	ProPhotoRGB:   RefIlluminantD50,
	SmptecRGB:     RefIlluminantD65,
	SRGB:          RefIlluminantD65,
	WideGamutRGB:  RefIlluminantD50,
}

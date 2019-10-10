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

package colorspace

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

const (
	GammaAdobeRGB      = 2.2
	GammaAppleRGB      = 1.8
	GammaBestRGB       = 2.2
	GammaBetaRGB       = 2.2
	GammaBruceRGB      = 2.2
	GammaCieRGB        = 2.2
	GammaColorMatchRGB = 1.8
	GammaDonRGB4       = 2.2
	GammaEktaSpacePS5  = 2.2
	GammaNtscRGB       = 2.2
	GammaPalSecamRGB   = 2.2
	GammaProPhotoRGB   = 1.8
	GammaSmptecRGB     = 2.2
	GammaSRGB          = 2.2
	GammaWideGamutRGB  = 2.2
)

var Gamma = map[string]float64{
	AdobeRGB:      GammaAdobeRGB,
	AppleRGB:      GammaAppleRGB,
	BestRGB:       GammaBestRGB,
	BetaRGB:       GammaBetaRGB,
	BruceRGB:      GammaBruceRGB,
	CieRGB:        GammaCieRGB,
	ColorMatchRGB: GammaColorMatchRGB,
	DonRGB4:       GammaDonRGB4,
	EktaSpacePS5:  GammaEktaSpacePS5,
	NtscRGB:       GammaNtscRGB,
	PalSecamRGB:   GammaPalSecamRGB,
	ProPhotoRGB:   GammaProPhotoRGB,
	SmptecRGB:     GammaSmptecRGB,
	SRGB:          GammaSRGB,
	WideGamutRGB:  GammaWideGamutRGB,
}

const (
	IlluminantAdobeRGB      = "D65"
	IlluminantAppleRGB      = "D65"
	IlluminantBestRGB       = "D50"
	IlluminantBetaRGB       = "D50"
	IlluminantBruceRGB      = "D65"
	IlluminantBT2020        = "D65"
	IlluminantCieRGB        = "E"
	IlluminantColorMatchRGB = "D50"
	IlluminantDonRGB4       = "D50"
	IlluminantEktaSpacePS5  = "D50"
	IlluminantNtscRGB       = "D50"
	IlluminantPalSecamRGB   = "D65"
	IlluminantProPhotoRGB   = "D50"
	IlluminantSmptecRGB     = "D65"
	IlluminantSRGB          = "D65"
	IlluminantWideGamutRGB  = "D50"
)

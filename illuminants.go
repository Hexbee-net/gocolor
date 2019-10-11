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

// Reference illuminants
const (
	RefIlluminantA         = "A"
	RefIlluminantB         = "B"
	RefIlluminantC         = "C"
	RefIlluminantD50       = "D50"
	RefIlluminantD55       = "D55"
	RefIlluminantD65       = "D65"
	RefIlluminantD75       = "D75"
	RefIlluminantE         = "E"
	RefIlluminantF2        = "F2"
	RefIlluminantF7        = "F7"
	RefIlluminantF11       = "F11"
	RefIlluminantBlackBody = "BlackBody"
)

// Standard observers
const (
	Observer2  = 2
	Observer10 = 10
)

const (
	CieE = 216.0 / 24389.0
	CieK = 24389.0 / 27.0
)

var observerWhitePoints = map[int]map[string]vector{
	Observer2: {
		RefIlluminantA:   vector{1.09850, 1.00000, 0.35585},
		RefIlluminantB:   vector{0.99072, 1.00000, 0.85223},
		RefIlluminantC:   vector{0.98074, 1.00000, 1.18232},
		RefIlluminantD50: vector{0.96422, 1.00000, 0.82521},
		RefIlluminantD55: vector{0.95682, 1.00000, 0.92149},
		RefIlluminantD65: vector{0.95047, 1.00000, 1.08883},
		RefIlluminantD75: vector{0.94972, 1.00000, 1.22638},
		RefIlluminantE:   vector{1.00000, 1.00000, 1.00000},
		RefIlluminantF2:  vector{0.99186, 1.00000, 0.67393},
		RefIlluminantF7:  vector{0.95041, 1.00000, 1.08747},
		RefIlluminantF11: vector{1.00962, 1.00000, 0.64350},
	},
	Observer10: {
		RefIlluminantD50: vector{0.96720, 1.00000, 0.8143},
		RefIlluminantD55: vector{0.95800, 1.00000, 0.9093},
		RefIlluminantD65: vector{0.94810, 1.00000, 1.0730},
		RefIlluminantD75: vector{0.94416, 1.00000, 1.2064},
	},
}

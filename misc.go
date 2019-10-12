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
)

// ApplyChromaticAdaptation applies a chromatic adaptation matrix to convert
// XYZ values between illuminants.
//
// It is important to recognize that color transformation results in color errors,
// determined by how far the original illuminant is from the target illuminant.
// For example, D65 to A could result in very high maximum deviance.
//
// Read https://web.stanford.edu/~sujason/ColorBalancing/adaptation.html and
// http://www.brucelindbloom.com/index.html?Eqn_ChromAdapt.html for more information
func ApplyChromaticAdaptation(x, y, z float64, source, target string, observer int, adaptation string) (float64, float64, float64) {
	// Get white-points for the observer
	obsWp, ok := observerWhitePoints[observer]
	if !ok {
		panic(fmt.Sprintf("unrecognized observer angle: %v", observer))
	}

	// Get white-points for illuminants
	srcWP, ok := obsWp[source]
	if !ok {
		panic(fmt.Sprintf("unrecognized source illuminant: %v", source))
	}
	tgtWP, ok := obsWp[target]
	if !ok {
		panic(fmt.Sprintf("unrecognized target illuminant: %v", target))
	}

	// Retrieve the appropriate transformation matrix from the constants.
	mTransform := getAdaptationMatrix(srcWP, tgtWP, adaptation)

	// Perform the adaptation.
	r := mTransform.vdot(vector{x, y, z})

	// Return individual X, Y, and Z coordinates.
	return r.v0, r.v1, r.v2
}

func getAdaptationMatrix(sourceWP, targetWP vector, adaptation string) matrix {
	// Get the appropriate transformation matrix.
	mAdp := chromaticAdaptation[adaptation]
	mInv := chromaticAdaptationInverse[adaptation]

	// Sharpened cone responses ~ rho gamma beta ~ sharpened r g b
	rgbSrc := mAdp.vdot(sourceWP)
	rgbTgt := mAdp.vdot(targetWP)

	// Ratio of whitepoint sharpened responses
	mRat := rgbTgt.vdiv(rgbSrc).diag()

	// Final transformation matrix
	return mInv.mdot(mRat).mdot(mAdp)
}

func getWhitePoint(observer int, illuminant string) vector {
	// Get white-points for the observer
	obsWp, ok := observerWhitePoints[observer]
	if !ok {
		panic(fmt.Sprintf("unrecognized observer angle: %v", observer))
	}

	// Get white-point for illuminant
	wp, ok := obsWp[illuminant]
	if !ok {
		panic(fmt.Sprintf("unrecognized illuminant: %v", illuminant))
	}

	return wp
}

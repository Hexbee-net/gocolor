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

type vector struct {
	v0, v1, v2 float64
}

func (v vector) vdiv(b vector) vector {
	return vector{
		v.v0 / b.v0,
		v.v1 / b.v1,
		v.v2 / b.v2,
	}
}

func (v vector) diag() matrix {
	return matrix{
		v.v0, 0, 0,
		0, v.v1, 0,
		0, 0, v.v2,
	}
}

func (v vector) mapfunc(f func(v float64) float64) vector {
	return vector{f(v.v0), f(v.v1), f(v.v2)}
}

////////////////////////////////////////

type matrix struct {
	m00, m01, m02 float64
	m10, m11, m12 float64
	m20, m21, m22 float64
}

func (a matrix) vdot(b vector) vector {
	return vector{
		a.m00*b.v0 + a.m01*b.v0 + a.m02*b.v0,
		a.m10*b.v1 + a.m11*b.v1 + a.m12*b.v1,
		a.m20*b.v2 + a.m21*b.v2 + a.m22*b.v2,
	}
}

func (a matrix) mdot(b matrix) matrix {
	return matrix{
		a.m00*b.m00 + a.m01*b.m10 + a.m02*b.m20,
		a.m00*b.m01 + a.m01*b.m11 + a.m02*b.m21,
		a.m00*b.m02 + a.m01*b.m12 + a.m02*b.m22,

		a.m10*b.m00 + a.m11*b.m10 + a.m12*b.m20,
		a.m10*b.m01 + a.m11*b.m11 + a.m12*b.m21,
		a.m10*b.m02 + a.m11*b.m12 + a.m12*b.m22,

		a.m20*b.m00 + a.m21*b.m10 + a.m22*b.m20,
		a.m20*b.m01 + a.m21*b.m11 + a.m22*b.m21,
		a.m20*b.m02 + a.m21*b.m12 + a.m22*b.m22,
	}
}

////////////////////////////////////////

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

func radians(v float64) float64 {
	return v * (math.Pi / 180)
}

func degrees(v float64) float64 {
	if v > 0 {
		return v * (180 / math.Pi)
	}

	return 360 - math.Abs(v)*(180/math.Pi)
}

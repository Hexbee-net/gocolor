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

type vector struct {
	v0, v1, v2 float64
}

func (a vector) vdiv(b vector) vector {
	return vector{
		a.v0 / b.v0,
		a.v1 / b.v1,
		a.v2 / b.v2,
	}
}

func (a vector) diag() matrix {
	return matrix{
		a.v0, 0, 0,
		0, a.v1, 0,
		0, 0, a.v2,
	}
}

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

func (a matrix) vdiv(b vector) matrix {
	return matrix{
		rgbTgt.v0 / rgbSrc.v0, 0, 0,
		0, rgbTgt.v1 / rgbSrc.v1, 0,
		0, 0, rgbTgt.v2 / rgbSrc.v2,
	}
}

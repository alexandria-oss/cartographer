// Copyright 2020 The Alexandria Foundation
//
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cartographer

// Axis Any axis in space
type Axis float64

// XAxis X point in space
type XAxis Axis

// YAxis Y point in space
type YAxis Axis

// ZAxis Z point in space
type ZAxis Axis

// XYCoordinate (x, y) point in a Two-Dimensional space
type XYCoordinate struct {
	X XAxis
	Y YAxis
}

// ToXYZCoordinate Get an XYZCoordinate reference
func (c XYCoordinate) ToXYZCoordinate(z ZAxis) *XYZCoordinate {
	// Z-axis default value
	lvl := ZAxis(0)

	if z != lvl {
		lvl = z
	}

	return &XYZCoordinate{
		X: c.X,
		Y: c.Y,
		Z: lvl,
	}
}

// XYZCoordinate (x,y,z) point in a Three-dimensional space
type XYZCoordinate struct {
	X XAxis
	Y YAxis
	Z ZAxis
}

// ToXYCoordinate Get XYCoordinate reference
func (c XYZCoordinate) ToXYCoordinate() *XYCoordinate {
	return &XYCoordinate{
		X: c.X,
		Y: c.Y,
	}
}

// FromXYCoordinate Convert and get a XYZCoordinate reference from a XYCoordinate
func (c XYZCoordinate) FromXYCoordinate(coordinate XYCoordinate) *XYZCoordinate {
	xDefault := XAxis(0)
	yDefault := YAxis(0)

	if coordinate.X != 0 {
		xDefault = coordinate.X
	}
	if coordinate.Y != 0 {
		yDefault = coordinate.Y
	}

	return &XYZCoordinate{
		X: xDefault,
		Y: yDefault,
		Z: c.Z,
	}
}

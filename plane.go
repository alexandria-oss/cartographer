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

// planeNodeList Plane data
type planeNodeList map[int64]*Node

// Plane Simple plane representation in space
//
// - One axis must be equal in all Vertex nodes
//
// - Minimum 3 Vertex nodes required to form a plane
type Plane struct {
	headVertex *Node
	nodes      planeNodeList
	totalSpace float64
}

// NewPlane Create a new plane
func NewPlane(n *Node) (*Plane, error) {
	p := &Plane{
		nodes:      make(planeNodeList),
		totalSpace: 0.0,
	}

	if err := p.SetBoundaries(n); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Plane) SetBoundaries(headVertex *Node) error {
	if headVertex == nil {
		return InvalidVertex
	}

	// Mathematical domain
	// 1. Vertex(k) Axis = Vertex(k+1) Axis
	// 2. Sum(vertex) >= 3
	vertexCount := 0
	nodes := headVertex.Iter()
	for nodes.Next() {
		vertexCount++
	}

	if vertexCount < 3 {
		return InvalidVertexRange
	}

	p.headVertex = headVertex
	return nil
}

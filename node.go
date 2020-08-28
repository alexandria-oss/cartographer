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

// Using Circular/Single Linked List as main data structure for Nodes

// Node is a key-value data type in Cartographer,
// it uses XYZCoordinate as key.
//
// - You may use any data type as value
//
// - A Node may be linked to another Node to create an inter-node data network
type Node struct {
	Key      *XYZCoordinate
	Value    interface{}
	Next     *Node
	iterator *Nodes
}

// Iter Start iteration over Node
func (n Node) Iter() *Nodes {
	// Singleton iterator
	if n.iterator == nil {
		n.iterator = &Nodes{currentValue: &n}
	}

	return n.iterator
}

// Nodes Iterator for Node
type Nodes struct {
	currentValue *Node
}

// Next Traverse to next Node
func (i *Nodes) Next() bool {
	if i.currentValue.Next != nil {
		i.currentValue = i.currentValue.Next
		return true
	}

	return false
}

// Value Get the current Node
func (i *Nodes) Value() *Node {
	return i.currentValue
}

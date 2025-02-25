/*
 * Copyright 2022 TikTok Pte. Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dag

import (
	"fmt"

	"github.com/tiktok/mia-rule-engine/inferencer/rete"
)

// RootNode the root node of virtualized Rete network
type RootNode struct {
	rete.RootNode
	id int64
}

// ID the identifier of root node
func (n *RootNode) ID() int64 {
	return n.id
}

// DOTID the display text of root node
func (n *RootNode) DOTID() string {
	return n.RootNode.String()
}

// TypeNode the type node of virtualized Rete network
type TypeNode struct {
	rete.TypeNode
	id int64
}

// ID the identifier of type node
func (n *TypeNode) ID() int64 {
	return n.id
}

// DOTID the display text of type node
func (n *TypeNode) DOTID() string {
	return n.TypeNode.String()
}

// AlphaNode the alpha node of virtualized Rete network
type AlphaNode struct {
	rete.AlphaNode
	id int64
}

// ID the identifier of alpha node
func (n *AlphaNode) ID() int64 {
	return n.id
}

// DOTID the display text of alpha node
func (n *AlphaNode) DOTID() string {
	return n.AlphaNode.String()
}

// BetaNode the beta node of virtualized Rete network
type BetaNode struct {
	rete.BetaNode
	id int64
}

// ID the identifier of beta node
func (n *BetaNode) ID() int64 {
	return n.id
}

// DOTID the display text of beta node
func (n *BetaNode) DOTID() string {
	return fmt.Sprintf("&(%d)", n.ID())
}

// TermNode the terminal node of virtualized Rete network
type TermNode struct {
	rete.TermNode
	id int64
}

// ID the identifier of terminal node
func (n *TermNode) ID() int64 {
	return n.id
}

// DOTID the display text of terminal node
func (n *TermNode) DOTID() string {
	return n.TermNode.String()
}

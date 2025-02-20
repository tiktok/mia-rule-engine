package dag

import (
	"fmt"

	"github.com/tiktok/mia-rule-engine/inferencer/rete"
)

type RootNode struct {
	rete.RootNode
	id int64
}

func (n *RootNode) ID() int64 {
	return n.id
}

func (n *RootNode) DOTID() string {
	return n.RootNode.String()
}

type TypeNode struct {
	rete.TypeNode
	id int64
}

func (n *TypeNode) ID() int64 {
	return n.id
}

func (n *TypeNode) DOTID() string {
	return n.TypeNode.String()
}

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

type AlphaNode struct {
	rete.AlphaNode
	id int64
}

func (n *AlphaNode) ID() int64 {
	return n.id
}

func (n *AlphaNode) DOTID() string {
	return n.AlphaNode.String()
}

type BetaNode struct {
	rete.BetaNode
	id int64
}

func (n *BetaNode) ID() int64 {
	return n.id
}

func (n *BetaNode) DOTID() string {
	return fmt.Sprintf("&(%d)", n.ID())
}

type TermNode struct {
	rete.TermNode
	id int64
}

func (n *TermNode) ID() int64 {
	return n.id
}

func (n *TermNode) DOTID() string {
	return n.TermNode.String()
}

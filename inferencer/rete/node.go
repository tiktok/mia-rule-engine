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

package rete

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/tiktok/mia-rule-engine/analyzer/rule"
)

type RootNode struct {
	id        string
	typeNodes map[string]*TypeNode
}

func (n *RootNode) String() string {
	return n.id
}

func (n *RootNode) GetTypes() map[string]*TypeNode {
	return n.typeNodes
}

type TypeNode struct {
	id         string
	lhs        string
	alphaNodes []*AlphaNode
}

func (n *TypeNode) String() string {
	return n.id
}

func (n *TypeNode) AlphaNodes() []*AlphaNode {
	return n.alphaNodes
}

type AlphaNode struct {
	id        string
	ops       string
	rhs       string
	rhsInt    *int64
	rhsFloat  *float64
	betaNodes []*BetaNode
	termNodes []*TermNode
}

func (n *AlphaNode) String() string {
	return fmt.Sprintf("%s %s", n.ops, n.rhs)
}

func (n *AlphaNode) BetaNodes() []*BetaNode {
	return n.betaNodes
}

func (n *AlphaNode) TermNodes() []*TermNode {
	return n.termNodes
}

func (n *AlphaNode) Compare(fact *Fact) bool {
	lhs := reflect.ValueOf(fact.val)

	if lhs.Kind().String() != "string" && n.isStr() {
		return false
	}

	switch lhs.Kind() {
	case reflect.String:
		return n.compareStr(lhs.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return n.compareInt(lhs.Int())
	case reflect.Float32, reflect.Float64:
		return n.compareFloat(lhs.Float())
	default:
		return false
	}
}

func (n *AlphaNode) isStr() bool {
	return n.rhsInt == nil && n.rhsFloat == nil
}

func (n *AlphaNode) compareStr(lhs string) bool {
	rhs := n.rhs
	rhs = strings.Trim(rhs, "\"")
	switch n.ops {
	case "=":
		return lhs == rhs
	case "!=":
		return lhs != rhs
	default:
		return false
	}
}

func (n *AlphaNode) compareInt(lhs int64) bool {
	if n.rhsInt == nil {
		return false
	}
	switch n.ops {
	case "=":
		return lhs == *n.rhsInt
	case "!=":
		return lhs != *n.rhsInt
	case ">":
		return lhs > *n.rhsInt
	case "<":
		return lhs < *n.rhsInt
	case ">=":
		return lhs >= *n.rhsInt
	case "<=":
		return lhs <= *n.rhsInt
	default:
		return false
	}
}

func (n *AlphaNode) compareFloat(lhs float64) bool {
	if n.rhsFloat == nil {
		return false
	}
	switch n.ops {
	case "=":
		return lhs == *n.rhsFloat
	case "!=":
		return lhs != *n.rhsFloat
	case ">":
		return lhs > *n.rhsFloat
	case "<":
		return lhs < *n.rhsFloat
	case ">=":
		return lhs >= *n.rhsFloat
	case "<=":
		return lhs <= *n.rhsFloat
	default:
		return false
	}
}

type BetaNode struct {
	id        string
	degree    int64
	termNodes []*TermNode
}

func (n *BetaNode) String() string {
	return n.id
}

func (n *BetaNode) TermNodes() []*TermNode {
	return n.termNodes
}

type TermNode struct {
	id   string
	rule *rule.Rule
}

func (n *TermNode) String() string {
	decision := n.rule.Decision()
	if n.rule.Priority() != nil {
		return fmt.Sprintf("%s priority: %s", decision.String(), n.rule.Priority().String())
	}
	return decision.String()
}

type Fact struct {
	key string
	val interface{}
}

func newFact(key string, val interface{}) *Fact {
	return &Fact{
		key: key,
		val: val,
	}
}

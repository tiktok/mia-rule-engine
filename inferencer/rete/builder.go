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
	"sort"
	"strconv"
	"strings"

	"github.com/tiktok/mia-rule-engine/analyzer/rule"
	cmpl "github.com/tiktok/mia-rule-engine/parser"
)

// Builder the struct of building Rete network
type Builder struct {
	root     RootNode
	typeMap  map[string]*TypeNode
	alphaMap map[string]*AlphaNode
	betaMap  map[string]*BetaNode
	termMap  map[string]*TermNode
}

// NewBuilder returns a new builder object
func NewBuilder() Builder {
	return Builder{
		root: RootNode{
			id:        "root",
			typeNodes: make(map[string]*TypeNode),
		},
		typeMap:  make(map[string]*TypeNode),
		alphaMap: make(map[string]*AlphaNode),
		betaMap:  make(map[string]*BetaNode),
		termMap:  make(map[string]*TermNode),
	}
}

// AddRule adds new rule into the builder
func (builder *Builder) AddRule(rule rule.Rule) {
	termNode := builder.getOrCreateTermNode(rule)
	for _, condition := range rule.Conditions() {
		// OR logic between these condition
		alphaNodes := make([]*AlphaNode, 0) // used for loop cache
		for _, expression := range condition.Expressions() {
			// AND logic between these expressions, so both Alpha nodes and Beta nodes will be added here
			typeNode, exist := builder.getOrCreateTypeNode(expression)
			if !exist {
				builder.root.typeNodes[typeNode.String()] = typeNode
			}
			alphaNode, exist := builder.getOrCreateAlphaNode(expression)
			if !exist {
				typeNode.alphaNodes = append(typeNode.alphaNodes, alphaNode)
			}
			alphaNodes = append(alphaNodes, alphaNode)
		}
		if len(condition.Expressions()) == 1 {
			// when only one expression, alpha node connects to terminal node directly for performance
			alphaNodes[0].termNodes = append(alphaNodes[0].termNodes, termNode)
		} else {
			betaNode, exist := builder.getOrCreateBetaNode(condition.Expressions())
			if !exist {
				for _, alphaNode := range alphaNodes {
					alphaNode.betaNodes = append(alphaNode.betaNodes, betaNode)
					betaNode.degree++
				}
			}
			betaNode.termNodes = append(betaNode.termNodes, termNode)
		}
	}
}

// Build return the built Rete network
func (builder *Builder) Build() Network {
	return Network{
		defaults: make([]rule.Decision, 0),
		root:     builder.root,
	}
}

func (builder *Builder) getOrCreateTypeNode(expression cmpl.IExpressionContext) (*TypeNode, bool) {
	id := expression.Lhs().GetText()
	typeNode, exist := builder.typeMap[id]
	if !exist {
		typeNode = &TypeNode{
			id:         id,
			lhs:        expression.Lhs().GetText(),
			alphaNodes: make([]*AlphaNode, 0),
		}
		builder.typeMap[id] = typeNode
	}
	return typeNode, exist
}

func (builder *Builder) getOrCreateAlphaNode(expression cmpl.IExpressionContext) (*AlphaNode, bool) {
	id := expression.GetText()
	alphaNode, exist := builder.alphaMap[id]
	if !exist {
		var (
			rhsInt   *int64
			rhsFloat *float64
		)
		rhs := expression.Rhs().GetText()
		intVal, err := strconv.ParseInt(rhs, 10, 64)
		if err == nil {
			rhsInt = &intVal
		}
		floatVal, err := strconv.ParseFloat(rhs, 64)
		if err == nil {
			rhsFloat = &floatVal
		}
		alphaNode = &AlphaNode{
			id:        id,
			ops:       expression.OPS().GetText(),
			rhs:       rhs,
			rhsInt:    rhsInt,
			rhsFloat:  rhsFloat,
			betaNodes: make([]*BetaNode, 0),
			termNodes: make([]*TermNode, 0),
		}
		builder.alphaMap[id] = alphaNode
	}
	return alphaNode, exist
}

func (builder *Builder) getOrCreateBetaNode(expressions []cmpl.IExpressionContext) (*BetaNode, bool) {
	exps := make([]string, len(expressions))
	for idx := range expressions {
		exps[idx] = expressions[idx].GetText()
	}
	sort.Strings(exps)
	id := strings.Join(exps, "|")
	betaNode, exist := builder.betaMap[id]
	if !exist {
		betaNode = &BetaNode{
			id:        id,
			degree:    0,
			termNodes: make([]*TermNode, 0),
		}
		builder.betaMap[id] = betaNode
	}
	return betaNode, exist
}

func (builder *Builder) getOrCreateTermNode(rule rule.Rule) *TermNode {
	id := rule.Ctx().GetText()
	termNode, exist := builder.termMap[id]
	if !exist {
		termNode = &TermNode{
			id:   id,
			rule: &rule,
		}
		builder.termMap[id] = termNode
	}
	return termNode
}

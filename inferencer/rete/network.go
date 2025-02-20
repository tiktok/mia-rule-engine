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
	"github.com/antlr4-go/antlr/v4"
	"github.com/tiktok/mia-rule-engine/analyzer"
	"github.com/tiktok/mia-rule-engine/analyzer/rule"
	cmpl "github.com/tiktok/mia-rule-engine/parser"
)

type Input struct {
	facts  []*Fact
	memory *Memory
}

func NewInput(fields []interface{}) *Input {
	facts := make([]*Fact, 0)
	for i := 0; i < len(fields); i = i + 2 {
		facts = append(facts, newFact(fields[i].(string), fields[i+1]))
	}
	return &Input{
		facts:  facts,
		memory: NewMemory(),
	}
}

// Network is read-only for concurrency-safe
type Network struct {
	defaults []rule.Decision
	root     RootNode
}

func BuildNetwork(policy string) Network {
	input := antlr.NewInputStream(policy)

	lexer := cmpl.NewcmplLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := cmpl.NewcmplParser(stream)

	parser.RemoveErrorListeners()
	parser.AddErrorListener(analyzer.NewErrorListener())

	parser.BuildParseTrees = true // validate the syntax the build the AST
	tree := parser.Policy()
	listener := analyzer.NewParser()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	builder := NewBuilder()
	for _, r := range listener.Rules() {
		builder.AddRule(r)
	}
	network := builder.Build()

	network.defaults = listener.Defaults()
	return network
}

func (nw *Network) Defaults() []rule.Decision {
	return nw.defaults
}

func (nw *Network) Root() RootNode {
	return nw.root
}

// Accept executes logics to prepare for the infer process
func (nw *Network) Accept(input *Input) {
	for _, decision := range nw.defaults {
		switch decision.Type() {
		case rule.Scope:
			input.memory.agenda.scopes[decision.Key()] = &decision
		case rule.Action:
			input.memory.agenda.actions[decision.Key()] = &decision
		case rule.Fact:
			input.memory.agenda.newFacts = append(input.memory.agenda.newFacts, newFact(decision.Key(), decision.Val()))
		}
	}
}

// Infer only supports one extra iteration of new fact to avoid performance issues
func (nw *Network) Infer(input *Input) (map[string]*rule.Decision, map[string]*rule.Decision, map[string]*rule.Rule) {
	agenda := nw.infer(input)
	if len(agenda.newFacts) > 0 {
		input.facts = agenda.newFacts
		agenda = nw.infer(input)
	}
	return agenda.scopes, agenda.actions, agenda.rules
}

func (nw *Network) infer(input *Input) *Agenda {
	for _, fact := range input.facts {
		if typeNode, matched := nw.root.GetTypes()[fact.key]; matched {
			for _, alphaNode := range typeNode.AlphaNodes() {
				if alphaNode.Compare(fact) {
					for _, betaNode := range alphaNode.BetaNodes() {
						betas := input.memory.state.betas
						if _, exist := betas[betaNode]; !exist {
							betas[betaNode] = 0
						}
						betas[betaNode]++
						if betas[betaNode] == betaNode.degree {
							nw.updateAgenda(input.memory.agenda, betaNode.TermNodes())
						}
					}
					nw.updateAgenda(input.memory.agenda, alphaNode.TermNodes())
				}
			}
		}
	}
	return input.memory.agenda
}

func (nw *Network) updateAgenda(agenda *Agenda, termNodes []*TermNode) {
	for _, termNode := range termNodes {
		decision := termNode.rule.Decision()
		if decision.Type() == rule.Fact {
			agenda.newFacts = append(agenda.newFacts, &Fact{
				key: decision.Key(),
				val: decision.Val(),
			})
		} else {
			if r, found := agenda.rules[decision.Key()]; found {
				if !r.Prioritize(termNode.rule) {
					continue
				}
			}
			switch decision.Type() {
			case rule.Scope:
				agenda.scopes[decision.Key()] = &decision
			case rule.Action:
				agenda.actions[decision.Key()] = &decision
			}
			agenda.rules[decision.Key()] = termNode.rule
		}
	}
}

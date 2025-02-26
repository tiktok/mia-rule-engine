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

	"github.com/tiktok/mia-rule-engine/analyzer/rule"
)

// Memory uses to store the intermediate states of inference process
type Memory struct {
	state  *State
	agenda *Agenda
}

// NewMemory returns a new memory object
func NewMemory() *Memory {
	return &Memory{
		state:  newState(),
		agenda: NewAgenda(),
	}
}

// Agenda stores the inference results for output return
type Agenda struct {
	scopes   map[string]*rule.Decision
	actions  map[string]*rule.Decision
	rules    map[string]*rule.Rule
	newFacts []*Fact // intermediate states
}

// NewAgenda returns a new agenda object
func NewAgenda() *Agenda {
	return &Agenda{
		scopes:   make(map[string]*rule.Decision),
		actions:  make(map[string]*rule.Decision),
		rules:    make(map[string]*rule.Rule),
		newFacts: make([]*Fact, 0),
	}
}

// String return the serialized agenda
func (a *Agenda) String() string {
	str := "[agenda]\n"
	for _, scope := range a.scopes {
		str = fmt.Sprintf("%s%s\n", str, scope.String())
	}
	for _, action := range a.actions {
		str = fmt.Sprintf("%s%s\n", str, action.String())
	}
	for _, r := range a.rules {
		str = fmt.Sprintf("%s%s\n", str, r.String())
	}
	return str
}

// State stores the Rete network inference states
type State struct {
	betas map[*BetaNode]int64
}

func newState() *State {
	return &State{
		betas: make(map[*BetaNode]int64),
	}
}

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

package engine

import (
	"github.com/tiktok/mia-rule-engine/analyzer/rule"
	"github.com/tiktok/mia-rule-engine/inferencer/rete"
)

// Engine the engine for inference.
type Engine struct {
	*rete.Network
}

// Result the output of the inference process.
type Result struct {
	Scopes   map[string]*rule.Decision
	Actions  map[string]*rule.Decision
	HitRules map[string]*rule.Rule
}

// New returns a new engine object based on the policy (rule set).
func New(policy string) *Engine {
	network := rete.BuildNetwork(policy)
	return &Engine{
		Network: &network,
	}
}

// Eval the key of result map will be the key of assignments, such as the account of account:banned.
// 1. provide benefit to fastly retrieve the targeted scope(s).
// 2. avoid duplicated keys of scope to mess the results.
func (e *Engine) Eval(fields ...interface{}) *Result {
	input := rete.NewInput(fields)
	e.Accept(input)
	scopes, actions, rules := e.Infer(input)
	return &Result{
		Scopes:   scopes,
		Actions:  actions,
		HitRules: rules,
	}
}

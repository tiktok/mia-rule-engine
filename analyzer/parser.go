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

package analyzer

import (
	"github.com/tiktok/mia-rule-engine/analyzer/rule"
	cmpl "github.com/tiktok/mia-rule-engine/parser"
)

// Parser the parser to parse ANTLR file and store the details
type Parser struct {
	*cmpl.BasecmplListener
	defaults []rule.Decision
	rules    []rule.Rule
}

// NewParser return the new parser object
func NewParser() *Parser {
	return &Parser{
		defaults: make([]rule.Decision, 0),
		rules:    make([]rule.Rule, 0),
	}
}

// EnterDefault handles the parsing logics of default statement
func (pl *Parser) EnterDefault(ctx *cmpl.DefaultContext) {
	pl.defaults = append(pl.defaults, rule.DecisionFromCtx(ctx.Decision()))
}

// EnterRule handles the parsing logics of rule statement
func (pl *Parser) EnterRule(ctx *cmpl.RuleContext) {
	pl.rules = append(pl.rules, rule.FromCtx(ctx))
}

// Defaults returns the default decisions
func (pl *Parser) Defaults() []rule.Decision {
	return pl.defaults
}

// Rules returns the rules.
func (pl *Parser) Rules() []rule.Rule {
	return pl.rules
}

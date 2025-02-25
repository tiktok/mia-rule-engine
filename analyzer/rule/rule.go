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

package rule

import (
	"fmt"

	cmpl "github.com/tiktok/mia-rule-engine/parser"
)

// Rule the deserialized rule for logic processing.
type Rule struct {
	// TODO: add more fields here after the framework is settled
	ctx        cmpl.IRuleContext
	decision   Decision
	conditions []Condition
	priority   *Priority
}

// FromCtx gets the rule from the parsed ANTLR file.
func FromCtx(ctx *cmpl.RuleContext) Rule {
	decision := DecisionFromCtx(ctx.Decision())
	conditions := ConditionsFromCtx(ctx.Conditions())
	r := Rule{
		ctx:        ctx,
		decision:   decision,
		conditions: conditions,
	}
	if ctx.Priority() != nil {
		p := PriorityFromCtx(ctx.Priority())
		r.priority = &p
	}
	return r
}

// Ctx returns the ANTLR rule context.
func (r *Rule) Ctx() cmpl.IRuleContext {
	return r.ctx
}

// Decision returns the decision of the rule.
func (r *Rule) Decision() Decision {
	return r.decision
}

// Conditions returns the conditions of the rule.
func (r *Rule) Conditions() []Condition {
	return r.conditions
}

// Priority returns the priority of the rule.
func (r *Rule) Priority() *Priority {
	return r.priority
}

// String returns the serialized rule.
func (r *Rule) String() string {
	str := fmt.Sprintf("[Rule]\nStatement:%s\nDecision:%s\nConditions:\n", r.ctx.GetText(), r.decision.String())
	for _, condition := range r.conditions {
		str = fmt.Sprintf("%s%s\n", str, condition.String())
	}
	if r.priority != nil {
		str = fmt.Sprintf("%sPriority:%s\n", str, r.priority.String())
	}
	return str
}

// Prioritize compares the priority between rules.
func (r *Rule) Prioritize(rr *Rule) bool {
	// Assumption: the rules in comparison has been validated with the same scope keys
	// 1. a rule without priority specified is the highest priority, to make it backward compatible
	// 2. if both two rules are without priority specified or with same priority, pick the original one
	// 3. always pick the higher priority one (the smaller num)
	if rr == nil {
		return false
	}
	if r.priority == nil {
		return false
	}
	if rr.priority == nil {
		return true
	}
	return r.priority.Val() > rr.priority.Val()
}

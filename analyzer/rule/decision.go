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

type Decision struct {
	assignment cmpl.IAssignmentContext
	t          DType
	key        string
	val        string
}

func DecisionFromCtx(ctx cmpl.IDecisionContext) Decision {
	return Decision{
		assignment: ctx.Assignment(),
		t:          DTypeFromCtx(ctx.Type_()),
		key:        ctx.Assignment().Key().GetText(),
		val:        ctx.Assignment().Val().GetText(),
	}
}

func (d *Decision) Assignment() cmpl.IAssignmentContext {
	return d.assignment
}

func (d *Decision) Type() DType {
	return d.t
}

func (d *Decision) String() string {
	return fmt.Sprintf("%s: %s", d.t.String(), d.assignment.GetText())
}

func (d *Decision) Key() string {
	return d.key
}

func (d *Decision) Val() string {
	return d.val
}

type DType int8

const (
	UnknownTypeStr = "[unknown]"
	ScopeStr       = "scope"
	ActionStr      = "action"
	FactStr        = "fact"
)

const (
	UnknownType DType = iota
	Scope
	Action
	Fact
)

func DTypeFromCtx(ctx cmpl.ITypeContext) DType {
	switch {
	case ctx.SCOPE() != nil:
		return Scope
	case ctx.ACTION() != nil:
		return Action
	case ctx.FACT() != nil:
		return Fact
	default:
		return UnknownType
	}
}

func (dt DType) String() string {
	switch dt {
	case Scope:
		return ScopeStr
	case Action:
		return ActionStr
	case Fact:
		return FactStr
	default:
		return UnknownTypeStr
	}
}

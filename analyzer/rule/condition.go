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
	"strings"

	cmpl "github.com/tiktok/mia-rule-engine/parser"
)

// Condition sets AND logic between all comparisons here, and it is OR logic between all conditions
type Condition struct {
	expressions []cmpl.IExpressionContext
}

func (c Condition) Expressions() []cmpl.IExpressionContext {
	return c.expressions
}

func (c Condition) String() string {
	expressions := make([]string, 0)
	for _, expression := range c.expressions {
		expressions = append(expressions, expression.GetText())
	}
	condition := strings.Join(expressions, " AND ")
	return fmt.Sprintf("[%s]", condition)
}

// ConditionsFromCtx flattens the OR logics into DNF (Disjunctive Normal Form）
// TODO: need to write full coverage UTs to ensure the logic correctness
func ConditionsFromCtx(ctx cmpl.IConditionsContext) []Condition {
	return flattenConditions(ctx)
}

// flattenConditions handles ASTs like
// - condition (AND|OR condition)*
func flattenConditions(ctx cmpl.IConditionsContext) []Condition {
	// matrix is introduced here to handling the comparison aggregation within parentheses
	conditionMatrix := make([][]Condition, 0)
	conditionMatrix = appendNewConditions(conditionMatrix)
	if len(ctx.AllCondition()) != 0 {
		for idx := range ctx.AllCondition() {
			conditions := flattenCondition(ctx.Condition(idx))
			// always merge the latest one
			mergeConditions(conditionMatrix, conditions)
			if ctx.Lops(idx) != nil && ctx.Lops(idx).OR() != nil {
				conditionMatrix = appendNewConditions(conditionMatrix)
			}
		}
	}
	return flattenMatrix(conditionMatrix)
}

// flattenCondition handles ASTs like
// - expression (AND|OR (condition))
// - expression (AND|OR expression)*
func flattenCondition(ctx cmpl.IConditionContext) []Condition {
	// matrix is introduced here to handling the comparison aggregation within parentheses
	conditionMatrix := make([][]Condition, 0)
	conditionMatrix = appendNewConditions(conditionMatrix)
	if len(ctx.AllCondition()) != 0 {
		// handle condition with parentheses
		// P.S. based on the current grammar, the nested condition can only happen in cases with parentheses
		conditions := flattenCondition(ctx.Condition(0))
		// always merge the latest one
		mergeConditions(conditionMatrix, conditions)
	} else {
		for idx := range ctx.AllExpression() {
			mergeExpressions(conditionMatrix, ctx.Expression(idx))
			if ctx.Lops(idx) != nil && ctx.Lops(idx).OR() != nil {
				conditionMatrix = appendNewConditions(conditionMatrix)
			}
		}
	}
	return flattenMatrix(conditionMatrix)
}

func flattenMatrix(conditionMatrix [][]Condition) []Condition {
	conditions := make([]Condition, 0)
	for i := range conditionMatrix {
		for j := range conditionMatrix[i] {
			conditions = append(conditions, conditionMatrix[i][j])
		}
	}
	return conditions
}

func appendNewConditions(conditionMatrix [][]Condition) [][]Condition {
	conditions := make([]Condition, 0)
	conditions = append(conditions, Condition{expressions: make([]cmpl.IExpressionContext, 0)})
	return append(conditionMatrix, conditions)
}

func mergeConditions(conditionMatrix [][]Condition, conditions []Condition) {
	if len(conditions) == 0 {
		return
	}

	combo := make([]Condition, len(conditionMatrix[len(conditionMatrix)-1])*len(conditions))
	idx := 0
	for i := range conditionMatrix[len(conditionMatrix)-1] {
		expressions := conditionMatrix[len(conditionMatrix)-1][i].expressions
		exps := make([]cmpl.IExpressionContext, len(expressions))
		copy(exps, expressions)
		for j := range conditions {
			combo[idx] = Condition{expressions: append(exps, conditions[j].expressions...)}
			idx++
		}
	}
	conditionMatrix[len(conditionMatrix)-1] = combo
}

func mergeExpressions(conditionMatrix [][]Condition, expression cmpl.IExpressionContext) {
	if expression == nil {
		return
	}
	condition := conditionMatrix[len(conditionMatrix)-1]
	for i := range condition {
		condition[i].expressions = append(condition[i].expressions, expression)
	}
}

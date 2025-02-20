// Code generated from parser/cmpl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // cmpl

import "github.com/antlr4-go/antlr/v4"

// cmplListener is a complete listener for a parse tree produced by cmplParser.
type cmplListener interface {
	antlr.ParseTreeListener

	// EnterPolicy is called when entering the policy production.
	EnterPolicy(c *PolicyContext)

	// EnterDefault is called when entering the default production.
	EnterDefault(c *DefaultContext)

	// EnterRule is called when entering the rule production.
	EnterRule(c *RuleContext)

	// EnterDecision is called when entering the decision production.
	EnterDecision(c *DecisionContext)

	// EnterItem is called when entering the item production.
	EnterItem(c *ItemContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterConditions is called when entering the conditions production.
	EnterConditions(c *ConditionsContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterLops is called when entering the lops production.
	EnterLops(c *LopsContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterPriority is called when entering the priority production.
	EnterPriority(c *PriorityContext)

	// EnterLhs is called when entering the lhs production.
	EnterLhs(c *LhsContext)

	// EnterRhs is called when entering the rhs production.
	EnterRhs(c *RhsContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterVal is called when entering the val production.
	EnterVal(c *ValContext)

	// ExitPolicy is called when exiting the policy production.
	ExitPolicy(c *PolicyContext)

	// ExitDefault is called when exiting the default production.
	ExitDefault(c *DefaultContext)

	// ExitRule is called when exiting the rule production.
	ExitRule(c *RuleContext)

	// ExitDecision is called when exiting the decision production.
	ExitDecision(c *DecisionContext)

	// ExitItem is called when exiting the item production.
	ExitItem(c *ItemContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitConditions is called when exiting the conditions production.
	ExitConditions(c *ConditionsContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitLops is called when exiting the lops production.
	ExitLops(c *LopsContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitPriority is called when exiting the priority production.
	ExitPriority(c *PriorityContext)

	// ExitLhs is called when exiting the lhs production.
	ExitLhs(c *LhsContext)

	// ExitRhs is called when exiting the rhs production.
	ExitRhs(c *RhsContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitVal is called when exiting the val production.
	ExitVal(c *ValContext)
}

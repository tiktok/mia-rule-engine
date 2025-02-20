// Code generated from parser/cmpl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // cmpl

import "github.com/antlr4-go/antlr/v4"

// BasecmplListener is a complete listener for a parse tree produced by cmplParser.
type BasecmplListener struct{}

var _ cmplListener = &BasecmplListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecmplListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecmplListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecmplListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecmplListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPolicy is called when production policy is entered.
func (s *BasecmplListener) EnterPolicy(ctx *PolicyContext) {}

// ExitPolicy is called when production policy is exited.
func (s *BasecmplListener) ExitPolicy(ctx *PolicyContext) {}

// EnterDefault is called when production default is entered.
func (s *BasecmplListener) EnterDefault(ctx *DefaultContext) {}

// ExitDefault is called when production default is exited.
func (s *BasecmplListener) ExitDefault(ctx *DefaultContext) {}

// EnterRule is called when production rule is entered.
func (s *BasecmplListener) EnterRule(ctx *RuleContext) {}

// ExitRule is called when production rule is exited.
func (s *BasecmplListener) ExitRule(ctx *RuleContext) {}

// EnterDecision is called when production decision is entered.
func (s *BasecmplListener) EnterDecision(ctx *DecisionContext) {}

// ExitDecision is called when production decision is exited.
func (s *BasecmplListener) ExitDecision(ctx *DecisionContext) {}

// EnterItem is called when production item is entered.
func (s *BasecmplListener) EnterItem(ctx *ItemContext) {}

// ExitItem is called when production item is exited.
func (s *BasecmplListener) ExitItem(ctx *ItemContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BasecmplListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BasecmplListener) ExitParameters(ctx *ParametersContext) {}

// EnterConditions is called when production conditions is entered.
func (s *BasecmplListener) EnterConditions(ctx *ConditionsContext) {}

// ExitConditions is called when production conditions is exited.
func (s *BasecmplListener) ExitConditions(ctx *ConditionsContext) {}

// EnterCondition is called when production condition is entered.
func (s *BasecmplListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BasecmplListener) ExitCondition(ctx *ConditionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasecmplListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasecmplListener) ExitExpression(ctx *ExpressionContext) {}

// EnterType is called when production type is entered.
func (s *BasecmplListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BasecmplListener) ExitType(ctx *TypeContext) {}

// EnterLops is called when production lops is entered.
func (s *BasecmplListener) EnterLops(ctx *LopsContext) {}

// ExitLops is called when production lops is exited.
func (s *BasecmplListener) ExitLops(ctx *LopsContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BasecmplListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BasecmplListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BasecmplListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BasecmplListener) ExitParameter(ctx *ParameterContext) {}

// EnterPriority is called when production priority is entered.
func (s *BasecmplListener) EnterPriority(ctx *PriorityContext) {}

// ExitPriority is called when production priority is exited.
func (s *BasecmplListener) ExitPriority(ctx *PriorityContext) {}

// EnterLhs is called when production lhs is entered.
func (s *BasecmplListener) EnterLhs(ctx *LhsContext) {}

// ExitLhs is called when production lhs is exited.
func (s *BasecmplListener) ExitLhs(ctx *LhsContext) {}

// EnterRhs is called when production rhs is entered.
func (s *BasecmplListener) EnterRhs(ctx *RhsContext) {}

// ExitRhs is called when production rhs is exited.
func (s *BasecmplListener) ExitRhs(ctx *RhsContext) {}

// EnterKey is called when production key is entered.
func (s *BasecmplListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BasecmplListener) ExitKey(ctx *KeyContext) {}

// EnterVal is called when production val is entered.
func (s *BasecmplListener) EnterVal(ctx *ValContext) {}

// ExitVal is called when production val is exited.
func (s *BasecmplListener) ExitVal(ctx *ValContext) {}

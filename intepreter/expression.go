package intepreter

import (
	"dklang/token"
)

// Expressions
type Exp interface {
	EvalExpression(env *Env) interface{}
}

type BinaryExpression struct {
	Left     Exp
	Operator token.Token
	Right    Exp
}

type ValueExpWrapper struct{ Value Value }

func (b BinaryExpression) EvalExpression(env *Env) interface{} {
	leftVal := b.Left.EvalExpression(env)
	rightVal := b.Right.EvalExpression(env)

	switch b.Operator.Type {
	case token.PLUS:
		leftNum, okLeft := leftVal.(int)
		rightNum, okRight := rightVal.(int)
		if !okLeft || !okRight {
			panic("Type error: Add between two integers only supported")
		}
		return leftNum + rightNum
	case token.MINUS:
		leftNum, okLeft := leftVal.(int)
		rightNum, okRight := rightVal.(int)
		if !okLeft || !okRight {
			panic("Type error: Add between two integers only supported")
		}
		return leftNum + rightNum
	default:
		panic("Not a valid binary operator")
	}

}

func (v ValueExpWrapper) EvalExpression(env *Env) interface{} {
	value := v.Value.EvalValue(env)
	return value
}

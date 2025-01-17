package intepreter

import (
	"fmt"
)

// Statements
type Statement interface {
	EvalStatement(env *Env) interface{}
}

type Body struct {
	Body []Statement
}

func (b Body) EvalStatement(env *Env) interface{} {
	for i := range len(b.Body) {
		switch value := b.Body[i].EvalStatement(env).(type) {
		case ReturnValue:
			return value
		}
	}
	return nil
}

type AssignStatement struct {
	Name  string
	Value Exp
}

func (a AssignStatement) EvalStatement(env *Env) interface{} {
	value := a.Value.EvalExpression(env)
	env.Variables[a.Name] = value
	return true
}

type ExpStatementWrapper struct{ Exp Exp }

func (e ExpStatementWrapper) EvalStatement(env *Env) interface{} {
	return e.Exp.EvalExpression(env)
}

type ReturnStatement struct{ Exp Exp }

func (r ReturnStatement) EvalStatement(env *Env) interface{} {
	if env.ParentEnv == nil {
		panic("Return statement in top level")
	}
	value := r.Exp.EvalExpression(env)
	return ReturnValue{Value: value}
}

type IfStatement struct {
	Condition Exp
	IfBody    Statement
	ElseBody  Statement
}

func (i IfStatement) EvalStatement(env *Env) interface{} {
	switch condition := i.Condition.EvalExpression(env).(type) {
	case bool:
		if condition {
			return i.IfBody.EvalStatement(env)
		} else {
			return i.ElseBody.EvalStatement(env)
		}
	default:
		err := fmt.Sprintf("Cannot determine condition of type %T", condition)
		panic(err)
	}
}

type WhileStatement struct {
	Condition Exp
	Body      Statement
}

func (w WhileStatement) EvalStatement(env *Env) interface{} {
loop:
	for {
		switch condition := w.Condition.EvalExpression(env).(type) {
		case bool:
			if !condition {
				break loop
			}
		default:
			err := fmt.Sprintf("Cannot determine condition of type %T", condition)
			panic(err)
		}
		switch value := w.Body.EvalStatement(env).(type) {
		case ReturnValue:
			return value
		}
	}
	return nil
}

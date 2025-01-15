package intepreter

import "fmt"

// Statements
type Statement interface {
	EvalStatement(env *Env) interface{}
}

type Body struct {
	Body []Statement
}

func (b Body) EvalStatement(env *Env) interface{} {
	for i := range len(b.Body) {
		val := b.Body[i].EvalStatement(env)
		fmt.Println(val)
	}
	return 0
}

type Assign struct {
	Name  string
	Value Exp
}

func (a Assign) EvalStatement(env *Env) interface{} {
	value := a.Value.EvalExpression(env)
	env.Variables[a.Name] = value
	return value
}

type ExpStatementWrapper struct{ Exp Exp }

func (e ExpStatementWrapper) EvalStatement(env *Env) interface{} {
	value := e.Exp.EvalExpression(env)
	return value
}

type ReturnStatement struct{ Exp Exp }

func (r ReturnStatement) EvalStatement(env *Env) interface{} {
	value := r.Exp.EvalExpression(env)
	return value
}

type FuncDef struct {
	Name       string
	Parameters []string
	Body       Statement
}

func (f FuncDef) EvalStatement(env *Env) interface{} {
	env.Functions[f.Name] = f
	return 0
}

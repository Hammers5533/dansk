package intepreter

// Statements
type Statement interface {
	EvalStatement(env *Env) interface{}
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

type FuncDef struct {
	Name       string
	Parameters []string
	Body       Exp
}

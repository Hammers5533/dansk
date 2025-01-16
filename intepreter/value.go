package intepreter

import "fmt"

// Values
type Value interface {
	EvalValue(env *Env) interface{}
}

type Float struct{ Value float64 }

func (f Float) EvalValue(env *Env) interface{} {
	return f.Value
}

type Integer struct{ Value int }

func (i Integer) EvalValue(env *Env) interface{} {
	return i.Value
}

type String struct{ Value string }

func (s String) EvalValue(env *Env) interface{} {
	return s.Value
}

type Variable struct{ Value string }

func (v Variable) EvalValue(env *Env) interface{} {
	val, ok := env.checkVariable(v.Value)

	if !ok {
		err := fmt.Sprintf("Identifier %s not defined\n", v.Value)
		panic(err)
	}
	return val
}

type Bool struct{ Value bool }

func (b Bool) EvalValue(env *Env) interface{} {
	return b.Value
}

type FuncDef struct {
	Name       string
	Parameters []string
	Body       Statement
}

func (f FuncDef) EvalValue(env *Env) interface{} {
	env.Variables[f.Name] = f
	return true
}

type InternalFunc struct {
	Name       string
	Parameters []string
	Func       func(...any) any
}

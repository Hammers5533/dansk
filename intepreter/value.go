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
	val, ok := env.Variables[v.Value]
	if !ok {
		s := fmt.Sprintf("Variable %s referenced before assignment", v.Value)
		panic(s)
	}
	return val
}

type FunctionName struct{ Value string }

func (n FunctionName) EvalValue(env *Env) interface{} {
	val, ok := env.Functions[n.Value]
	if !ok {
		s := fmt.Sprintf("Function %s called before assignment", n.Value)
		panic(s)
	}
	return val
}

type Bool struct{ Value bool }

func (b Bool) EvalValue(env *Env) interface{} {
	return b.Value
}

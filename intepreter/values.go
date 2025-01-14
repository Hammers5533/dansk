package intepreter

import "fmt"

// Values
type Value interface {
	EvalValue(env *Env) interface{}
}

type Float struct{ Value float32 }

type Integer struct{ Value int }

func (i Integer) EvalValue(env *Env) interface{} {
	return i.Value
}

type String struct{ Value string }

type Variable struct{ Value string }

func (v Variable) EvalValue(env *Env) interface{} {
	val, ok := env.Variables[v.Value]
	if !ok {
		s := fmt.Sprintf("Variable %s referenced before assignment", v.Value)
		panic(s)
	}
	return val
}

type Env struct {
	Variables map[string]interface{}
	Functions map[string]FuncDef
}

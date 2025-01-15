package intepreter

import (
	"fmt"
)

type Program struct {
	Body []Statement
}

func (p Program) Run() int {
	env := &Env{
		Variables: make(map[string]interface{}),
		Functions: make(map[string]FuncDef),
	}
	for _, statement := range p.Body {
		switch statement.(type) {
		case Body:
			panic("Cannot parse body inside another body")
		default:
			val := statement.EvalStatement(env)
			fmt.Println(val)
		}
	}
	return 0
}

type Env struct {
	Variables map[string]interface{}
	Functions map[string]FuncDef
}

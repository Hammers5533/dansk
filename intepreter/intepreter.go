package intepreter

import "fmt"

type Program struct {
	Statements []Statement
}

func (p Program) Run() int {
	env := &Env{
		Variables: make(map[string]interface{}),
		Functions: make(map[string]FuncDef),
	}
	for i := range len(p.Statements) {
		val := p.Statements[i].EvalStatement(env)
		fmt.Println(val)
	}
	return 0
}

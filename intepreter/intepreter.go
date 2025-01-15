package intepreter

type Program struct {
	Body Statement
}

func (p Program) Run() int {
	env := &Env{
		Variables: make(map[string]interface{}),
		Functions: make(map[string]FuncDef),
	}
	p.Body.EvalStatement(env)
	return 0
}

type Env struct {
	Variables map[string]interface{}
	Functions map[string]FuncDef
}

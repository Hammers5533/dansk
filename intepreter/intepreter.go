package intepreter

type Program struct {
	Body Body
}

func (p Program) Run() int {
	env := &Env{
		ParentEnv: nil,
		Variables: make(map[string]interface{}),
	}
	env.Variables["meddel"] = builtinPrint
	p.Body.EvalStatement(env)
	return 0
}

type Env struct {
	ParentEnv *Env
	Variables map[string]interface{}
}

func (env *Env) checkVariable(name string) (interface{}, bool) {
	val, ok := env.Variables[name]
	if ok {
		return val, true
	}
	if env.ParentEnv == nil {
		return nil, false
	}
	return env.ParentEnv.checkVariable(name)
}

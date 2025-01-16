package intepreter

import (
	"fmt"
)

var builtinPrint InternalFunc = InternalFunc{
	Name:       "meddel",
	Parameters: []string{"printval"},
	Func: func(a ...any) any {
		fmt.Println(a...)
		return nil
	},
}

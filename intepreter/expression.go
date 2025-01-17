package intepreter

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Hammers5533/dklang/token"
)

// Expressions
type Exp interface {
	EvalExpression(env *Env) interface{}
}

type BinaryExpression struct {
	Left     Exp
	Operator token.Token
	Right    Exp
}

type ValueExpWrapper struct{ Value Value }

func (b BinaryExpression) EvalExpression(env *Env) interface{} {
	leftVal := b.Left.EvalExpression(env)
	rightVal := b.Right.EvalExpression(env)

	if !checkTypes(leftVal, rightVal) {
		err := fmt.Sprintf("Cannot perform binary operator on types %T and %T", leftVal, rightVal)
		panic(err)
	}

	switch value := leftVal.(type) {
	case int:
		switch b.Operator.Type {
		case token.PLUS:
			return value + rightVal.(int)
		case token.MINUS:
			return value - rightVal.(int)
		case token.MULTIPLY:
			return value * rightVal.(int)
		case token.DIVIDE:
			return value / rightVal.(int)
		case token.MODULUS:
			return value % rightVal.(int)
		case token.LEQ:
			return value <= rightVal.(int)
		case token.GEQ:
			return value >= rightVal.(int)
		case token.LT:
			return value < rightVal.(int)
		case token.GT:
			return value > rightVal.(int)
		case token.NEQ:
			return value != rightVal.(int)
		case token.EQ:
			return value == rightVal.(int)
		default:
			err := fmt.Sprintf("Undefined operator for type %T", leftVal)
			panic(err)
		}
	case float64:
		switch b.Operator.Type {
		case token.PLUS:
			return value + rightVal.(float64)
		case token.MINUS:
			return value - rightVal.(float64)
		case token.MULTIPLY:
			return value * rightVal.(float64)
		case token.DIVIDE:
			return value / rightVal.(float64)
		case token.MODULUS:
			err := fmt.Sprintf("Modulus operator not supported for type %T", leftVal)
			panic(err)
		case token.LEQ:
			return value <= rightVal.(float64)
		case token.GEQ:
			return value >= rightVal.(float64)
		case token.LT:
			return value < rightVal.(float64)
		case token.GT:
			return value > rightVal.(float64)
		case token.NEQ:
			return value != rightVal.(float64)
		case token.EQ:
			return value == rightVal.(float64)
		default:
			err := fmt.Sprintf("Undefined operator for type %T", leftVal)
			panic(err)
		}
	case bool:
		switch b.Operator.Type {
		case token.AND:
			return value && rightVal.(bool)
		case token.OR:
			return value || rightVal.(bool)
		default:
			err := fmt.Sprintf("Undefined operator for type %T", leftVal)
			panic(err)
		}
	default:
		err := fmt.Sprintf("Type %T not valid for binary operator", leftVal)
		panic(err)
	}
}

func (v ValueExpWrapper) EvalExpression(env *Env) interface{} {
	value := v.Value.EvalValue(env)
	return value
}

type FuncCall struct {
	Name       Exp
	Parameters []Exp
}

func (c FuncCall) EvalExpression(env *Env) interface{} {
	funcVariables := make(map[string]interface{})

	switch funcDef := c.Name.EvalExpression(env).(type) {
	case FuncDef:
		if len(funcDef.Parameters) != len(c.Parameters) {
			err := fmt.Sprintf("Input parameters for %s does not match required parameters %s", c.Name, strings.Join(funcDef.Parameters, ", "))
			panic(err)
		}

		// Assign expressions to variables
		for i := range len(funcDef.Parameters) {
			parameterExp := c.Parameters[i].EvalExpression(env)
			funcVariables[funcDef.Parameters[i]] = parameterExp
		}

		functionEnv := &Env{
			ParentEnv: env,
			Variables: funcVariables,
		}

		switch value := funcDef.Body.EvalStatement(functionEnv).(type) {
		case ReturnValue:
			return value.Value
		default:
			err := fmt.Sprintf("Expected return value for function call but got %T", value)
			panic(err)
		}
	case InternalFunc:
		if len(funcDef.Parameters) != len(c.Parameters) {
			err := fmt.Sprintf("Input parameters for %s does not match required parameters %s", c.Name, strings.Join(funcDef.Parameters, ", "))
			panic(err)
		}

		parameters := []any{}
		for i := range len(funcDef.Parameters) {
			parameterExp := c.Parameters[i].EvalExpression(env)
			parameters = append(parameters, parameterExp)
		}

		return funcDef.Func(parameters...)
	default:
		err := fmt.Sprintf("Type %T is not a callable method", funcDef)
		panic(err)
	}
}

type AssignExpression struct {
	Name  Exp
	Value Exp
}

func (a AssignExpression) EvalExpression(env *Env) interface{} {

	switch Name := a.Name.(type) {
	case ValueExpWrapper:
		switch Wrapper := Name.Value.(type) {
		case Variable:
			_, ok := env.Variables[Wrapper.Value]
			if !ok {
				err := fmt.Sprintf("Cannot assign value to indefined variable %s", Wrapper.Value)
				panic(err)
			}
			env.Variables[Wrapper.Value] = a.Value.EvalExpression(env)
		default:
			panic("Cannot Assign value to non-identifier")
		}
	default:
		panic("Cannot assign value to expression")
	}
	return true
}

func checkTypes(left, right interface{}) bool {
	return reflect.TypeOf(left) == reflect.TypeOf(right)
}

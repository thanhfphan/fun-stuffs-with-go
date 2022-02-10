package be

import (
	"fmt"
)

var (
	NULL         = &Null{}
	TRUE_OBJECT  = &Boolean{Value: true}
	FALSE_OBJECT = &Boolean{Value: false}
)

// Eval ...
func Eval(node Node, env *Environment) Object {
	switch node := node.(type) {
	case *Program:
		return evalProgram(node, env)
	case *ExpressionStatement:
		return Eval(node.Expression, env)
	case *ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if val == nil || val.Type() == ERROR_OBJ {
			return val
		}
		return &ReturnValue{Value: val}
	case *IntegerLiteral:
		return &Integer{Value: node.Value}
	case *StringLiteral:
		return &String{Value: node.Value}
	case *BooleanLiteral:
		return nativeBoolToBooleanObject(node.Value)
	case *PrefixExpression:
		right := Eval(node.Right, env)
		if right == nil || right.Type() == ERROR_OBJ {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *InfixExpression:
		left := Eval(node.Left, env)
		if left == nil || left.Type() == ERROR_OBJ {
			return left
		}

		right := Eval(node.Right, env)
		if right == nil || right.Type() == ERROR_OBJ {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)
	case *Identifier:
		return evalIdentifier(node, env)
	}

	return nil
}

func nativeBoolToBooleanObject(input bool) *Boolean {
	if input {
		return TRUE_OBJECT
	}

	return FALSE_OBJECT
}

func evalProgram(program *Program, env *Environment) Object {
	var result Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *ReturnValue:
			return result.Value
		case *Error:
			return result
		}
	}

	return result
}

func evalPrefixExpression(operator string, right Object) Object {
	switch operator {
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return &Error{Message: fmt.Sprintf("unknown operator: %s%s", operator, right.Type())}
	}
}

func evalMinusPrefixOperatorExpression(right Object) Object {
	if right.Type() != INTEGER_OBJ {
		return &Error{Message: fmt.Sprintf("unknown operator: -%s", right.Type())}
	}

	value := right.(*Integer).Value
	return &Integer{Value: -value}
}

func evalInfixExpression(operator string, left, right Object) Object {
	switch {
	case left.Type() == INTEGER_OBJ && right.Type() == INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == STRING_OBJ && right.Type() == STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	case left.Type() != right.Type():
		return &Error{Message: fmt.Sprintf("type mismatch: %s %s %s", left.Type(), operator, right.Type())}
	default:
		return &Error{Message: fmt.Sprintf("unknown operator: %s %s %s", left.Type(), operator, right.Type())}
	}
}

func evalIntegerInfixExpression(operator string, left, right Object) Object {
	leftVal := left.(*Integer).Value
	rightVal := right.(*Integer).Value

	switch operator {
	case "+":
		return &Integer{Value: leftVal + rightVal}
	case "-":
		return &Integer{Value: leftVal - rightVal}
	case "*":
		return &Integer{Value: leftVal * rightVal}
	case "/":
		return &Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return &Error{Message: fmt.Sprintf("unknown operator: %s %s %s", left.Type(), operator, right.Type())}
	}
}

func evalStringInfixExpression(operator string, left, right Object) Object {
	if operator != "+" {
		return &Error{Message: fmt.Sprintf("unknow operator: %s %s %s", left.ToString(), operator, right.Type())}
	}
	leftVal := left.(*String).Value
	rightVal := right.(*String).Value

	return &String{Value: leftVal + rightVal}
}

func evalIdentifier(node *Identifier, env *Environment) Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	return &Error{"identifier not found: " + node.Value}
}

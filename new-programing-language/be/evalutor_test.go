package be

import "testing"

func TestEvalSimple(t *testing.T) {
	input := `(long > 5) and 3 == 3`
	lexer := NewLexer(input)
	parser := NewParser(lexer)
	program := parser.ParseProgram()
	env := NewEnvironment()
	env.Set("long1", &Integer{Value: 13})

	result := Eval(program, env)
	a, ok := result.(*Error)
	t.Log(a.Message)
	t.Log(a)
	t.Log(ok)
	// t.Log(result.ToString())
	// t.Log(result.Type())
	t.Fail()
}

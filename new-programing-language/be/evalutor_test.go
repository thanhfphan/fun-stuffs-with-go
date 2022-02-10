package be

import "testing"

func TestEvalSimple(t *testing.T) {
	input := `(long > 5) and 3 == 3`
	lexer := NewLexer(input)
	parser := NewParser(lexer)
	program := parser.ParseProgram()
	env := NewEnvironment()
	env.Set("long", &Integer{Value: 13})

	result := Eval(program, env)

	t.Log(result.ToString())
	t.Log(result.Type())
	t.Fail()
}

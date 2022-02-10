package main

import (
	"fmt"

	"thanhfphan.com/be"
)

func main() {
	input := `(long > 5) and 3 == 3`
	lexer := be.NewLexer(input)
	parser := be.NewParser(lexer)
	program := parser.ParseProgram()
	env := be.NewEnvironment()
	env.Set("long", &be.Integer{Value: 13})

	result := be.Eval(program, env)

	fmt.Println(result.ToString())
	fmt.Println(result.Type())
}

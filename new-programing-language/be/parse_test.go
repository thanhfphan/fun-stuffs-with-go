package be

import "testing"

func TestDemo(t *testing.T) {
	t.Skip()
	input := `(BUSINESS_TYPE == "kinh doanh") and (CAR_TYPE == "pickup") and (SEAT_NUMBER < 6)`
	lexer := NewLexer(input)
	parser := NewParser(lexer)
	program := parser.ParseProgram()

	t.Fatalf("%s", program.ToString())
}

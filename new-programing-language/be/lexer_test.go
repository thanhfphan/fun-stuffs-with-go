package be

import "testing"

func TestNextToken(t *testing.T) {
	input := `(BUSINESS_TYPE == "kinh doanh") and (CAR_TYPE == "pickup") and (SEAT_NUMBER < 6)`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LEFT_PAREN, "("},
		{IDENTIFIER, "BUSINESS_TYPE"},
		{EQUAL, "=="},
		{STRING, "kinh doanh"},
		{RIGHT_PAREN, ")"},
		{AND, "and"},
		{LEFT_PAREN, "("},
		{IDENTIFIER, "CAR_TYPE"},
		{EQUAL, "=="},
		{STRING, "pickup"},
		{RIGHT_PAREN, ")"},
		{AND, "and"},
		{LEFT_PAREN, "("},
		{IDENTIFIER, "SEAT_NUMBER"},
		{LESS_THAN, "<"},
		{INTEGER, "6"},
		{RIGHT_PAREN, ")"},
	}

	l := NewLexer(input)
	for index, tc := range tests {
		token := l.NextToken()

		if token.Type != tc.expectedType {
			t.Fatalf("test[%d] - tokentype failed, expected=%q, got=%q", index, tc.expectedType, token.Type)
		}
		if token.Literal != tc.expectedLiteral {
			t.Fatalf("test[%d] -token token literal failed, expected=%q, got=%q", index, tc.expectedLiteral, token.Literal)
		}
		t.Logf("type=%q, literal=%q", token.Type, token.Literal)
	}
}

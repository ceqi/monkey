package lexer

import (
	"ceqi/monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		//...
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, wanted: %q, got: %q)", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong, wanted: %q, got: %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestSkipWhitespace(t *testing.T) {
	l := New(`

	let the_game_begain`)

	l.eatWhitespace()

	got := l.ch

	if got != 'l' {
		t.Fatalf("l.skipWhitespace(), l.ch = %v, expected: l", got)
	}

}

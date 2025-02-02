package lexer

import (
	"testing"

	"github.com/adityathebe/monkey/token"
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
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	//
	// tests := []struct {
	// 	expectedType    token.TokenType
	// 	expectedLiteral string
	// }{
	// 	{token.ASSIGN, "="},
	// 	{token.PLUS, "+"},
	// 	{token.LPAREN, "("},
	// 	{token.RPAREN, ")"},
	// 	{token.LBRACE, "{"},
	// 	{token.RBRACE, "}"},
	// 	{token.COMMA, ","},
	// 	{token.SEMICOLON, ";"},
	// 	{token.EOF, ""},
	// }

	l := New(input)

	for i, td := range tests {
		tok := l.NextToken()

		if tok.Type != td.expectedType {
			t.Errorf("[%d] - tokenType wrong. expected=%q, got=%q", i, td.expectedType, tok.Type)
		}

		if tok.Literal != td.expectedLiteral {
			t.Errorf("[%d] - literal wrong. expected=%q, got=%q", i, td.expectedLiteral, tok.Literal)
		}

	}
}

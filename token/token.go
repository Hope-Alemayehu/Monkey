package token

import "strings"

//defined the tokentype to be string
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	//means the token we don't know about
	ILLEGAL = "ILLEGAL"
	//end of file: tells the parser later on that it can stop
	EOF = "EOF"

	//Identifiers + literals
	IDENT = "IDENT" // such as add, foobar, x, y
	INT   = "INT"   //123456789

	//operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	EQ       = "=="
	NOT_EQ   = "!="

	LT = "<"
	GT = ">"

	//Delimiters

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

//checks the keywords table to see whether the given identifier is a keyword
//if it is it returns a keyword's tokenType constant
//if not we get back the token.IDENT

//?????? why are we returning IDENT
func LookupIdent(ident string) TokenType {
	//could cause problems later
	ident = strings.ToLower(ident) // Ensure case insensitivity
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

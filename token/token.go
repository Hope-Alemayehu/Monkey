package token

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
	ASSIGN = "="
	PLUS   = "+"

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
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

//checks the keywords table to see whether the given identifier is a keyword
//if it is it returns a keyword's tokenType constant
//if not we get back the token.IDENT

//?????? why are we returning IDENT
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

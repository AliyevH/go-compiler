package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// NOT_EQ
const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	EQ        = "=="
	NOT_EQ    = "!="    //NOT_EQ
	IDENT     = "IDENT" // add, foobar, x, y, ...
	INT       = "INT"   // 1234567890
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	COMMA     = ","
	SEMICOLON = ";"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

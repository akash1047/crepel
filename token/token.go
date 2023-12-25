package token

type TokenType string

type Token struct {
	Type     TokenType
	Literal  string
	Position int
}

func (t *Token) Span() (int, int) {
	return t.Position, t.Position + len(t.Literal)
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"

	CHARACTER_CONST = "CHARACTER_CONST"
	INTEGER_CONST   = "INTEGER_CONST"
	FLOAT_CONST     = "FLOAT_CONST"
	STRING_CONST    = "STRING_CONST"

	// operators

	ASSIGN      = "="
	PLUS        = "+"
	MINUS       = "-"
	ASTERISK    = "*"
	SLASH       = "/"
	MODULO      = "%"
	BANG        = "!"
	TILDE       = "~"
	AND_CHAR    = "&"
	PIPE        = "|"
	PLUS_PLUS   = "++"
	MINUS_MINUS = "--"
	DOT         = "."
	ARROW       = "->"
	LSHIFT      = "<<"
	RSHIFT      = ">>"
	XOR         = "^"
	AND         = "&&"
	OR          = "||"
	QUESTION    = "?"
	EQ          = "=="
	NEQ         = "!="
	LT          = "<"
	GT          = ">"
	LEQ         = "<="
	GEQ         = ">="

	PLUS_ASSIGN   = "+="
	MINUS_ASSIGN  = "-="
	MUL_ASSIGN    = "*="
	DIV_ASSIGN    = "/="
	MODULO_ASSIGN = "%="

	LSHIFT_ASSIGN = "<<="
	RSHFIT_ASSIGN = ">>="

	BITAND_ASSIGN = "&="
	BITOR_ASSIGN  = "|="
	XOR_ASSIGN    = "^="

	COMMA = ","

	// delimeters

	COLON     = ":"
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LSQB   = "["
	RSQB   = "]"

	// special

	ELLIPSE = "..."

	HASH = "#"

	// preprocessors

	P_IF    = "#if"
	P_ELIF  = "#elif"
	P_ELSE  = "#else"
	P_ENDIF = "#endif"

	P_IFDEF  = "#ifdef"
	P_IFNDEF = "#ifndef"
	P_DEFINE = "#define"
	P_UNDEF  = "#undef"

	P_INCLUDE = "#include"
	P_LINE    = "#line"
	P_ERROR   = "#error"
	P_PRAGMA  = "#pragma"

	P_DEFINED = "#defined"

	// keywords [ excluding c23, underscore prefixed ]

	AUTO = "auto"

	BREAK = "break"

	CASE     = "case"
	CHAR     = "char"
	CONST    = "const"
	CONTINUE = "continue"

	DEFAULT = "default"
	DO      = "do"
	DOUBLE  = "double"

	ELSE   = "else"
	ENUM   = "enum"
	EXTERN = "extern"

	FLOAT = "float"
	FOR   = "for"

	GOTO = "goto"

	IF     = "if"
	INLINE = "inline"
	INT    = "int"

	LONG = "long"

	REGISTER = "register"
	RESTRICT = "restrict"
	RETURN   = "return"

	SHORT  = "short"
	SIGNED = "signed"
	SIZEOF = "sizeof"
	STATIC = "static"
	STRUCT = "struct"
	SWITCH = "switch"

	TYPEDEF = "typedef"

	UNION    = "union"
	UNSIGNED = "unsigned"

	VOID     = "void"
	VOLATILE = "volatile"

	WHILE = "while"
)

var keywords = map[string]TokenType{
	"auto": AUTO,

	"break": BREAK,

	"case":     CASE,
	"char":     CHAR,
	"const":    CONST,
	"continue": CONTINUE,

	"default": DEFAULT,
	"do":      DO,
	"double":  DOUBLE,

	"else":   ELSE,
	"enum":   ENUM,
	"extern": EXTERN,

	"float": FLOAT,
	"for":   FOR,

	"goto": GOTO,

	"if":     IF,
	"inline": INLINE,
	"int":    INT,

	"long": LONG,

	"register": REGISTER,
	"restrict": RESTRICT,
	"return":   RETURN,

	"short":  SHORT,
	"signed": SIGNED,
	"sizeof": SIZEOF,
	"static": STATIC,
	"struct": STRUCT,
	"switch": SWITCH,

	"typedef": TYPEDEF,

	"union":    UNION,
	"unsigned": UNSIGNED,

	"void":     VOID,
	"volatile": VOLATILE,

	"while": WHILE,
}

func LookupIdent(literal string) TokenType {
	if ty, ok := keywords[literal]; ok {
		return ty
	}

	return IDENT
}

var preprocessors = map[string]TokenType{
	"#if":    P_IF,
	"#elif":  P_ELIF,
	"#else":  P_ELSE,
	"#endif": P_ENDIF,

	"#ifdef":  P_IFDEF,
	"#ifndef": P_IFNDEF,
	"#define": P_DEFINE,
	"#undef":  P_UNDEF,

	"#include": P_INCLUDE,
	"#line":    P_LINE,
	"#error":   P_ERROR,
	"#pragma":  P_PRAGMA,

	"#defined": P_DEFINED,
}

func LookupPreprocessor(literal string) TokenType {
	if pre, ok := preprocessors[literal]; ok {
		return pre
	}

	// we could return ILLEGAL_PREPROCESSOR
	return ILLEGAL
}

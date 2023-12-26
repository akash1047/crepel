package lexer

type Error struct {
	Message string
	Line    string
	Span    [2]int
}

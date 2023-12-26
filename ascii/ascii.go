package ascii

func IsAlpha(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func IsDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func IsBinaryDigit(ch byte) bool {
	return ch == '0' || ch == '1'
}

func IsHexDigit(ch byte) bool {
	return IsDigit(ch) || ch >= 'a' && ch <= 'f' || ch >= 'A' && ch <= 'F'
}

func IsOctalDigit(ch byte) bool {
	return ch >= '0' && ch <= '7'
}

func IsLetter(ch byte) bool {
	return IsAlpha(ch) || IsLetter(ch) || ch == '_' || ch == '$'
}

package url

func isASCIIDigit(c byte) bool {
	return c >= 48 && c <= 57
}

func isASCIILetter(c byte) bool {
	return c >= 97 && c <= 122 || c >= 65 && c <= 90
}

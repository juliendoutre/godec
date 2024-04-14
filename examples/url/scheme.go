package url

import (
	"fmt"

	"github.com/juliendoutre/godec"
)

type Scheme struct {
	scheme *string
}

func (s Scheme) Encode() ([]byte, error) {
	for _, c := range []byte(*s.scheme) {
		if !(c == '+') && !(c == '.') && !(c == '-') && !isASCIIDigit(c) && !isASCIILetter(c) {
			return nil, fmt.Errorf("invalid character %q", c)
		}
	}

	return []byte(*s.scheme), nil
}

func (s Scheme) Decode(input []byte) ([]byte, error) {
	// Reject empty schemes.
	if len(input) == 0 {
		return nil, fmt.Errorf("expected a scheme")
	}

	// See https://datatracker.ietf.org/doc/html/rfc1738#section-2.1:
	// Scheme names consist of a sequence of characters. The lower case
	// letters "a"--"z", digits, and the characters plus ("+"), period
	// ("."), and hyphen ("-") are allowed. For resiliency, programs
	// interpreting URLs should treat upper case letters as equivalent to
	// lower case in scheme names (e.g., allow "HTTP" as well as "http").

	i := 0
	for i = 0; i < len(input); i++ {
		if !(input[i] == '+') && !(input[i] == '.') && !(input[i] == '-') && !isASCIIDigit(input[i]) && !isASCIILetter(input[i]) {
			break
		}
	}

	if i == 0 {
		return nil, fmt.Errorf("expected a scheme")
	}

	*s.scheme = string(input[:i])

	return input[i:], nil
}

var _ godec.Codec = Scheme{}

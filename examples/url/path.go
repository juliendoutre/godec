package url

import (
	"bytes"

	"github.com/juliendoutre/godec"
)

type Path struct {
	path *string
}

func (p Path) Encode() ([]byte, error) {
	return []byte(*p.path), nil
}

func (p Path) Decode(input []byte) ([]byte, error) {
	queryParamsStartIndex := bytes.IndexByte(input, '?')
	if queryParamsStartIndex != -1 {
		*p.path = string(input[:queryParamsStartIndex])
		return input[queryParamsStartIndex:], nil
	}

	*p.path = string(input[:])
	return input, nil
}

var _ godec.Codec = Path{}

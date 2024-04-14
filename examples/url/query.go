package url

import (
	"bytes"

	"github.com/juliendoutre/godec"
)

type Query struct {
	query *string
}

func (q Query) Encode() ([]byte, error) {
	if q.query != nil {
		return append([]byte("?"), []byte(*q.query)...), nil
	}

	return nil, nil
}

func (q Query) Decode(input []byte) ([]byte, error) {
	fragmentStartIndex := bytes.IndexByte(input, '#')
	if fragmentStartIndex == -1 {
		fragmentStartIndex = len(input)
	}

	if len(input) > 1 && input[0] == '?' {
		*q.query = string(input[1:fragmentStartIndex])
		return input[fragmentStartIndex:], nil
	}

	return input, nil
}

var _ godec.Codec = Query{}

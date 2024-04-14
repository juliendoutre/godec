package url

import "github.com/juliendoutre/godec"

type Fragment struct {
	fragment *string
}

func (f Fragment) Encode() ([]byte, error) {
	if f.fragment != nil {
		return append([]byte("#"), []byte(*f.fragment)...), nil
	}

	return nil, nil
}

func (f Fragment) Decode(input []byte) ([]byte, error) {
	if len(input) > 1 && input[0] == '#' {
		*f.fragment = string(input[1:])
	}

	return nil, nil
}

var _ godec.Codec = Fragment{}

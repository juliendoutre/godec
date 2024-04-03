package godec

import (
	"bytes"
	"fmt"
	"strconv"
)

type ExactMatch []byte

func (e ExactMatch) Encode() ([]byte, error) {
	return e, nil
}

func (e ExactMatch) Decode(input []byte) ([]byte, error) {
	if bytes.Equal(e, input[:len(e)]) {
		return input[len(e):], nil
	}

	return nil, fmt.Errorf("expected %q", string(e))
}

var _ Codec = ExactMatch{}

type HexadecimalUInt8 struct {
	Variable *uint8
}

func (h HexadecimalUInt8) Encode() ([]byte, error) {
	return []byte(fmt.Sprintf("%02x", uint64(*h.Variable))), nil
}

func (h HexadecimalUInt8) Decode(input []byte) ([]byte, error) {
	if len(input) < 2 {
		return nil, fmt.Errorf("expected at least 2 bytes")
	}

	n, err := strconv.ParseUint(string(input[:2]), 16, 8)
	if err != nil {
		return nil, err
	}

	*h.Variable = uint8(n)

	return input[2:], nil
}

var _ Codec = HexadecimalUInt8{}

type NoMoreBytes struct{}

func (n NoMoreBytes) Encode() ([]byte, error) {
	return nil, nil
}

func (n NoMoreBytes) Decode(input []byte) ([]byte, error) {
	if len(input) != 0 {
		return nil, fmt.Errorf("expected no more bytes")
	}

	return input, nil
}

var _ Codec = NoMoreBytes{}

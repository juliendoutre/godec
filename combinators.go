package godec

type Sequence []Codec

func (s Sequence) Encode() ([]byte, error) {
	var out []byte

	for _, codec := range s {
		codecOut, err := codec.Encode()
		if err != nil {
			return nil, err
		}

		out = append(out, codecOut...)
	}

	return out, nil
}

func (s Sequence) Decode(input []byte) (remaining []byte, err error) {
	remaining = input

	for _, codec := range s {
		remaining, err = codec.Decode(remaining)
		if err != nil {
			return nil, err
		}
	}

	return remaining, nil
}

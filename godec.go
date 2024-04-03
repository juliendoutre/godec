package godec

type Encoder interface {
	Encode() ([]byte, error)
}

type Decoder interface {
	Decode(input []byte) ([]byte, error)
}

type Codec interface {
	Encoder
	Decoder
}

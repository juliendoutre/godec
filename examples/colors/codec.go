package colors

import "github.com/juliendoutre/godec"

func Codec(red, green, blue *uint8) godec.Sequence {
	return godec.Sequence([]godec.Codec{
		godec.ExactMatch([]byte("#")),
		godec.HexadecimalUInt8{Variable: red},
		godec.HexadecimalUInt8{Variable: green},
		godec.HexadecimalUInt8{Variable: blue},
		godec.NoMoreBytes{},
	})
}

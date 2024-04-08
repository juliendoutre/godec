package colors_test

import (
	"testing"
	"testing/quick"

	"github.com/juliendoutre/godec/examples/colors"
	"github.com/stretchr/testify/assert"
)

func TestColorEncoding(t *testing.T) {
	red := uint8(13)
	green := uint8(89)
	blue := uint8(42)
	codec := colors.Codec(&red, &green, &blue)

	out, err := codec.Encode()
	assert.NoError(t, err)
	assert.Equal(t, []byte("#0d592a"), out)
}

func TestValidColorDecoding(t *testing.T) {
	var red, green, blue uint8
	codec := colors.Codec(&red, &green, &blue)

	remainder, err := codec.Decode([]byte("#3399ff"))
	assert.NoError(t, err)
	assert.Equal(t, 0, len(remainder))
	assert.Equal(t, uint8(51), red)
	assert.Equal(t, uint8(153), green)
	assert.Equal(t, uint8(255), blue)
}

func TestInversibleProperty(t *testing.T) {
	f := func(expectedRed, expectedGreen, expectedBlue uint8) bool {
		actualRed := expectedRed
		actualGreen := expectedGreen
		actualBlue := expectedBlue

		encoder := colors.Codec(&expectedRed, &expectedGreen, &expectedBlue)
		out, err := encoder.Encode()
		if err != nil {
			return false
		}

		decoder := colors.Codec(&actualRed, &actualGreen, &actualBlue)
		remainder, err := decoder.Decode(out)
		if err != nil {
			return false
		}

		return len(remainder) == 0 && expectedRed == actualRed && expectedGreen == actualGreen && expectedBlue == actualBlue
	}

	if err := quick.Check(f, &quick.Config{}); err != nil {
		t.Error(err)
	}
}

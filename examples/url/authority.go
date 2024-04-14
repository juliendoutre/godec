package url

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/juliendoutre/godec"
)

type Authority struct {
	username *string
	password *string
	host     *string
	port     *uint
}

func (a Authority) Encode() ([]byte, error) {
	if a.host != nil {
		buffer := []byte("//")

		if a.username != nil {
			buffer = append(buffer, []byte(*a.username)...)
		}

		if a.password != nil {
			buffer = append(buffer, []byte(":")...)
			buffer = append(buffer, []byte(*a.password)...)
		}

		if a.username != nil {
			buffer = append(buffer, []byte("@")...)
		}

		buffer = append(buffer, []byte(*a.host)...)

		if a.port != nil {
			buffer = append(buffer, []byte(fmt.Sprintf("%d", *a.port))...)
		}

		return buffer, nil
	}

	return nil, nil
}

func (a Authority) Decode(input []byte) ([]byte, error) {
	if len(input) < 2 || string(input[:2]) != "//" {
		return input, nil
	}

	hostStartIndex := 2

	userinfoSeparatorIndex := bytes.IndexByte(input[2:], '@')
	if userinfoSeparatorIndex != -1 {
		hostStartIndex = userinfoSeparatorIndex + 3

		passwordSeparatorIndex := bytes.IndexByte(input[2:userinfoSeparatorIndex+2], ':')
		if passwordSeparatorIndex != -1 {
			*a.username = string(input[2 : passwordSeparatorIndex+2])
			*a.password = string(input[passwordSeparatorIndex+3 : userinfoSeparatorIndex+2])
		} else {
			*a.username = string(input[2 : userinfoSeparatorIndex+2])
		}
	}

	pathSeparatorIndex := bytes.IndexByte(input[hostStartIndex:], '/')
	if pathSeparatorIndex == -1 {
		pathSeparatorIndex = len(input)
	}

	portSeparatorIndex := bytes.IndexByte(input[hostStartIndex:], ':')
	if portSeparatorIndex != -1 {
		*a.host = string(input[hostStartIndex : portSeparatorIndex+hostStartIndex])

		port, err := strconv.ParseUint(string(input[portSeparatorIndex+hostStartIndex+1:pathSeparatorIndex+hostStartIndex]), 10, 32)
		if err != nil {
			return nil, err
		}

		*a.port = uint(port)
	} else {
		*a.host = string(input[hostStartIndex : pathSeparatorIndex+hostStartIndex])
	}

	return input[pathSeparatorIndex+hostStartIndex:], nil
}

var _ godec.Codec = Authority{}

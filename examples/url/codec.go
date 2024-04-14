package url

import (
	"github.com/juliendoutre/godec"
)

func Codec(scheme, username, password, host *string, port *uint, path, query, fragment *string) godec.Sequence {
	return godec.Sequence([]godec.Codec{
		Scheme{scheme},
		godec.ExactMatch([]byte(":")),
		Authority{username, password, host, port},
		Path{path},
		Query{query},
		Fragment{fragment},
	})
}

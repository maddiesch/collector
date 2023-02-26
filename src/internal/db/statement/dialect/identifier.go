package dialect

import (
	"strconv"
)

func Identifier(n string) string {
	return strconv.Quote(n)
}

func IdentifierMap(n string, _ int) string {
	return Identifier(n)
}

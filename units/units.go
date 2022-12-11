package units

import (
	"fmt"
	"math"
)

type Unit string

const (
	None  = Unit("")
	Bytes = Unit("bytes")
	Bits  = Unit("bits")
)

var unitShort = map[Unit]string{
	Bytes: "B",
	Bits:  "b",
}

func (u Unit) Short() string {
	return unitShort[u]
}

func FmtFloat64(x float64, u Unit) string {
	x, p := SIPrefix(x)
	if p == "" && math.Round(x) == x {
		return fmt.Sprintf("%d%s", int64(x), u)
	}
	return fmt.Sprintf("%.2f%s%s", x, p, u)
}

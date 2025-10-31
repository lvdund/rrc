package ies

import "rrc/utils"

// SRS-ConfigAp-v1310-cyclicShiftAp-v1310 ::= ENUMERATED
type SrsConfigapV1310CyclicshiftapV1310 struct {
	Value utils.ENUMERATED
}

const (
	SrsConfigapV1310CyclicshiftapV1310EnumeratedNothing = iota
	SrsConfigapV1310CyclicshiftapV1310EnumeratedCs8
	SrsConfigapV1310CyclicshiftapV1310EnumeratedCs9
	SrsConfigapV1310CyclicshiftapV1310EnumeratedCs10
	SrsConfigapV1310CyclicshiftapV1310EnumeratedCs11
)

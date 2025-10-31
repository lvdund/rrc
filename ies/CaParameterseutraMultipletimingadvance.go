package ies

import "rrc/utils"

// CA-ParametersEUTRA-multipleTimingAdvance ::= ENUMERATED
type CaParameterseutraMultipletimingadvance struct {
	Value utils.ENUMERATED
}

const (
	CaParameterseutraMultipletimingadvanceEnumeratedNothing = iota
	CaParameterseutraMultipletimingadvanceEnumeratedSupported
)

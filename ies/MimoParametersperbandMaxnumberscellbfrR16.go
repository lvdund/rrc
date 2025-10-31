package ies

import "rrc/utils"

// MIMO-ParametersPerBand-maxNumberSCellBFR-r16 ::= ENUMERATED
type MimoParametersperbandMaxnumberscellbfrR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMaxnumberscellbfrR16EnumeratedNothing = iota
	MimoParametersperbandMaxnumberscellbfrR16EnumeratedN1
	MimoParametersperbandMaxnumberscellbfrR16EnumeratedN2
	MimoParametersperbandMaxnumberscellbfrR16EnumeratedN4
	MimoParametersperbandMaxnumberscellbfrR16EnumeratedN8
)

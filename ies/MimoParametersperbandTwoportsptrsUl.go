package ies

import "rrc/utils"

// MIMO-ParametersPerBand-twoPortsPTRS-UL ::= ENUMERATED
type MimoParametersperbandTwoportsptrsUl struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandTwoportsptrsUlEnumeratedNothing = iota
	MimoParametersperbandTwoportsptrsUlEnumeratedSupported
)

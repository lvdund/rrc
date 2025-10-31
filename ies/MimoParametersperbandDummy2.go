package ies

import "rrc/utils"

// MIMO-ParametersPerBand-dummy2 ::= ENUMERATED
type MimoParametersperbandDummy2 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandDummy2EnumeratedNothing = iota
	MimoParametersperbandDummy2EnumeratedSupported
)

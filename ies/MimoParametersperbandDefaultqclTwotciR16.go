package ies

import "rrc/utils"

// MIMO-ParametersPerBand-defaultQCL-TwoTCI-r16 ::= ENUMERATED
type MimoParametersperbandDefaultqclTwotciR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandDefaultqclTwotciR16EnumeratedNothing = iota
	MimoParametersperbandDefaultqclTwotciR16EnumeratedSupported
)

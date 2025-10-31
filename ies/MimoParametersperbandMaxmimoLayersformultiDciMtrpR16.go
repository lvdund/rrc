package ies

import "rrc/utils"

// MIMO-ParametersPerBand-maxMIMO-LayersForMulti-DCI-mTRP-r16 ::= ENUMERATED
type MimoParametersperbandMaxmimoLayersformultiDciMtrpR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMaxmimoLayersformultiDciMtrpR16EnumeratedNothing = iota
	MimoParametersperbandMaxmimoLayersformultiDciMtrpR16EnumeratedSupported
)

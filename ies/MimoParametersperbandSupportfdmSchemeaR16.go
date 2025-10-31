package ies

import "rrc/utils"

// MIMO-ParametersPerBand-supportFDM-SchemeA-r16 ::= ENUMERATED
type MimoParametersperbandSupportfdmSchemeaR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandSupportfdmSchemeaR16EnumeratedNothing = iota
	MimoParametersperbandSupportfdmSchemeaR16EnumeratedSupported
)

package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-CSI-N-Max2-r17 ::= ENUMERATED
type MimoParametersperbandMtrpCsiNMax2R17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpCsiNMax2R17EnumeratedNothing = iota
	MimoParametersperbandMtrpCsiNMax2R17EnumeratedSupported
)

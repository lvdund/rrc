package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-CSI-CMR-r17 ::= ENUMERATED
type MimoParametersperbandMtrpCsiCmrR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpCsiCmrR17EnumeratedNothing = iota
	MimoParametersperbandMtrpCsiCmrR17EnumeratedSupported
)

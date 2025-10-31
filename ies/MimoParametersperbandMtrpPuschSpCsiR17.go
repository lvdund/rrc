package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PUSCH-SP-CSI-r17 ::= ENUMERATED
type MimoParametersperbandMtrpPuschSpCsiR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpPuschSpCsiR17EnumeratedNothing = iota
	MimoParametersperbandMtrpPuschSpCsiR17EnumeratedSupported
)

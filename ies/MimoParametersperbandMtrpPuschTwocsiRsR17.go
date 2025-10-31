package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PUSCH-twoCSI-RS-r17 ::= ENUMERATED
type MimoParametersperbandMtrpPuschTwocsiRsR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpPuschTwocsiRsR17EnumeratedNothing = iota
	MimoParametersperbandMtrpPuschTwocsiRsR17EnumeratedSupported
)

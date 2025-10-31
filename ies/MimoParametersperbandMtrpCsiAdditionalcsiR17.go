package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-CSI-additionalCSI-r17 ::= ENUMERATED
type MimoParametersperbandMtrpCsiAdditionalcsiR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpCsiAdditionalcsiR17EnumeratedNothing = iota
	MimoParametersperbandMtrpCsiAdditionalcsiR17EnumeratedX1
	MimoParametersperbandMtrpCsiAdditionalcsiR17EnumeratedX2
)

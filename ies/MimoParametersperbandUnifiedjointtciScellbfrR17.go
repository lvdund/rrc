package ies

import "rrc/utils"

// MIMO-ParametersPerBand-unifiedJointTCI-SCellBFR-r17 ::= ENUMERATED
type MimoParametersperbandUnifiedjointtciScellbfrR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandUnifiedjointtciScellbfrR17EnumeratedNothing = iota
	MimoParametersperbandUnifiedjointtciScellbfrR17EnumeratedSupported
)

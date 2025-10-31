package ies

import "rrc/utils"

// MIMO-ParametersPerBand-unifiedJointTCI-Legacy-r17 ::= ENUMERATED
type MimoParametersperbandUnifiedjointtciLegacyR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandUnifiedjointtciLegacyR17EnumeratedNothing = iota
	MimoParametersperbandUnifiedjointtciLegacyR17EnumeratedSupported
)

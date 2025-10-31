package ies

import "rrc/utils"

// MIMO-ParametersPerBand-unifiedJointTCI-Legacy-SRS-r17 ::= ENUMERATED
type MimoParametersperbandUnifiedjointtciLegacySrsR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandUnifiedjointtciLegacySrsR17EnumeratedNothing = iota
	MimoParametersperbandUnifiedjointtciLegacySrsR17EnumeratedSupported
)

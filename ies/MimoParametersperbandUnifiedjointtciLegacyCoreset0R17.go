package ies

import "rrc/utils"

// MIMO-ParametersPerBand-unifiedJointTCI-Legacy-CORESET0-r17 ::= ENUMERATED
type MimoParametersperbandUnifiedjointtciLegacyCoreset0R17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandUnifiedjointtciLegacyCoreset0R17EnumeratedNothing = iota
	MimoParametersperbandUnifiedjointtciLegacyCoreset0R17EnumeratedSupported
)

package ies

import "rrc/utils"

// MIMO-ParametersPerBand-unifiedJointTCI-commonMultiCC-r17 ::= ENUMERATED
type MimoParametersperbandUnifiedjointtciCommonmulticcR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandUnifiedjointtciCommonmulticcR17EnumeratedNothing = iota
	MimoParametersperbandUnifiedjointtciCommonmulticcR17EnumeratedSupported
)

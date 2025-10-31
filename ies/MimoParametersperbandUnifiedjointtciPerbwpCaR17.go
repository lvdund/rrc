package ies

import "rrc/utils"

// MIMO-ParametersPerBand-unifiedJointTCI-perBWP-CA-r17 ::= ENUMERATED
type MimoParametersperbandUnifiedjointtciPerbwpCaR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandUnifiedjointtciPerbwpCaR17EnumeratedNothing = iota
	MimoParametersperbandUnifiedjointtciPerbwpCaR17EnumeratedSupported
)

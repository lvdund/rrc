package ies

import "rrc/utils"

// MIMO-ParametersPerBand-unifiedJointTCI-BeamAlignDLRS-r17 ::= ENUMERATED
type MimoParametersperbandUnifiedjointtciBeamaligndlrsR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandUnifiedjointtciBeamaligndlrsR17EnumeratedNothing = iota
	MimoParametersperbandUnifiedjointtciBeamaligndlrsR17EnumeratedSupported
)

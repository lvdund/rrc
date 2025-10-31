package ies

import "rrc/utils"

// MIMO-ParametersPerBand-unifiedJointTCI-PC-association-r17 ::= ENUMERATED
type MimoParametersperbandUnifiedjointtciPcAssociationR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandUnifiedjointtciPcAssociationR17EnumeratedNothing = iota
	MimoParametersperbandUnifiedjointtciPcAssociationR17EnumeratedSupported
)

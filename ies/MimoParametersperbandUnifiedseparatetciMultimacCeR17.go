package ies

import "rrc/utils"

// MIMO-ParametersPerBand-unifiedSeparateTCI-multiMAC-CE-r17 ::= SEQUENCE
type MimoParametersperbandUnifiedseparatetciMultimacCeR17 struct {
	MinbeamapplicationtimeR17 MimoParametersperbandUnifiedseparatetciMultimacCeR17MinbeamapplicationtimeR17
	MaxactivateddlTciperccR17 utils.INTEGER `lb:0,ub:8`
	MaxactivatedulTciperccR17 utils.INTEGER `lb:0,ub:8`
}

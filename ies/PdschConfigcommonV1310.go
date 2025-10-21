package ies

import "rrc/utils"

// PDSCH-ConfigCommon-v1310 ::= SEQUENCE
type PdschConfigcommonV1310 struct {
	PdschMaxnumrepetitioncemodeaR13 *utils.ENUMERATED
	PdschMaxnumrepetitioncemodebR13 *utils.ENUMERATED
}

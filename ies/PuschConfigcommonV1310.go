package ies

import "rrc/utils"

// PUSCH-ConfigCommon-v1310 ::= SEQUENCE
type PuschConfigcommonV1310 struct {
	PuschMaxnumrepetitioncemodeaR13 *utils.ENUMERATED
	PuschMaxnumrepetitioncemodebR13 *utils.ENUMERATED
	PuschHoppingoffsetV1310         *utils.INTEGER
}

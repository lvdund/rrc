package ies

import "rrc/utils"

// PUSCH-ConfigCommon-v1310 ::= SEQUENCE
type PuschConfigcommonV1310 struct {
	PuschMaxnumrepetitioncemodeaR13 *PuschConfigcommonV1310PuschMaxnumrepetitioncemodeaR13
	PuschMaxnumrepetitioncemodebR13 *PuschConfigcommonV1310PuschMaxnumrepetitioncemodebR13
	PuschHoppingoffsetV1310         *utils.INTEGER `lb:0,ub:maxAvailNarrowBandsR13`
}

package ies

import "rrc/utils"

// MBS-RNTI-SpecificConfigId-r17 ::= utils.INTEGER (0..maxG-RNTI-1-r17)
type MbsRntiSpecificconfigidR17 struct {
	Value utils.INTEGER `lb:0,ub:maxGRnti1R17`
}

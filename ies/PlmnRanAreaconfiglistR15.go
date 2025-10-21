package ies

import "rrc/utils"

// PLMN-RAN-AreaConfigList-r15 ::= SEQUENCE OF PLMN-RAN-AreaConfig-r15
// SIZE (1..maxPLMN-r15)
type PlmnRanAreaconfiglistR15 struct {
	Value utils.Sequence[PlmnRanAreaconfigR15]
}

package ies

import "rrc/utils"

// MPE-ResourceId-r17 ::= utils.INTEGER (1..maxMPE-Resources-r17)
type MpeResourceidR17 struct {
	Value utils.INTEGER `lb:0,ub:maxMPEResourcesR17`
}

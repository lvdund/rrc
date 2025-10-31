package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-v1530-ce-PUSCH-FlexibleStartPRB-AllocConfig-r15-setup ::= SEQUENCE
type PuschConfigdedicatedV1530CePuschFlexiblestartprbAllocconfigR15Setup struct {
	OffsetceModebR15 *utils.INTEGER `lb:0,ub:3`
}

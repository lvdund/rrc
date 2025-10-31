package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-v1530-ce-PUSCH-SubPRB-Config-r15-setup ::= SEQUENCE
type PuschConfigdedicatedV1530CePuschSubprbConfigR15Setup struct {
	LocationceModebR15      *utils.INTEGER `lb:0,ub:5`
	SixtonecyclicshiftR15   utils.INTEGER  `lb:0,ub:3`
	ThreetonecyclicshiftR15 utils.INTEGER  `lb:0,ub:2`
}

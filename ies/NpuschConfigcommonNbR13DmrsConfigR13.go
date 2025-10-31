package ies

import "rrc/utils"

// NPUSCH-ConfigCommon-NB-r13-dmrs-Config-r13 ::= SEQUENCE
type NpuschConfigcommonNbR13DmrsConfigR13 struct {
	ThreetoneBasesequenceR13  *utils.INTEGER `lb:0,ub:12`
	ThreetoneCyclicshiftR13   utils.INTEGER  `lb:0,ub:2`
	SixtoneBasesequenceR13    *utils.INTEGER `lb:0,ub:14`
	SixtoneCyclicshiftR13     utils.INTEGER  `lb:0,ub:3`
	TwelvetoneBasesequenceR13 *utils.INTEGER `lb:0,ub:30`
}

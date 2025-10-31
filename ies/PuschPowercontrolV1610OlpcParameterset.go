package ies

import "rrc/utils"

// PUSCH-PowerControl-v1610-olpc-ParameterSet ::= SEQUENCE
type PuschPowercontrolV1610OlpcParameterset struct {
	OlpcParametersetdci01R16 *utils.INTEGER `lb:0,ub:2`
	OlpcParametersetdci02R16 *utils.INTEGER `lb:0,ub:2`
}

package ies

import "rrc/utils"

// CrossCarrierSchedulingConfig-carrierIndicatorSize-r16 ::= SEQUENCE
type CrosscarrierschedulingconfigCarrierindicatorsizeR16 struct {
	Carrierindicatorsizedci12R16 utils.INTEGER `lb:0,ub:3`
	Carrierindicatorsizedci02R16 utils.INTEGER `lb:0,ub:3`
}

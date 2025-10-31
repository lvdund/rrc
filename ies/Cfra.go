package ies

import "rrc/utils"

// CFRA ::= SEQUENCE
// Extensible
type Cfra struct {
	Occasions                *CfraOccasions
	Resources                CfraResources
	TotalnumberofraPreambles *utils.INTEGER `lb:0,ub:63,ext`
}

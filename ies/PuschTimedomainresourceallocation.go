package ies

import "rrc/utils"

// PUSCH-TimeDomainResourceAllocation ::= SEQUENCE
type PuschTimedomainresourceallocation struct {
	K2                   *utils.INTEGER `lb:0,ub:32`
	Mappingtype          PuschTimedomainresourceallocationMappingtype
	Startsymbolandlength utils.INTEGER `lb:0,ub:127`
}

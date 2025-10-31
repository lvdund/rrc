package ies

import "rrc/utils"

// PDSCH-TimeDomainResourceAllocation ::= SEQUENCE
type PdschTimedomainresourceallocation struct {
	K0                   *utils.INTEGER `lb:0,ub:32`
	Mappingtype          PdschTimedomainresourceallocationMappingtype
	Startsymbolandlength utils.INTEGER `lb:0,ub:127`
}

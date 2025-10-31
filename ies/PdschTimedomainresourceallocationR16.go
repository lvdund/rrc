package ies

import "rrc/utils"

// PDSCH-TimeDomainResourceAllocation-r16 ::= SEQUENCE
// Extensible
type PdschTimedomainresourceallocationR16 struct {
	K0R16                   *utils.INTEGER `lb:0,ub:32`
	MappingtypeR16          PdschTimedomainresourceallocationR16MappingtypeR16
	StartsymbolandlengthR16 utils.INTEGER `lb:0,ub:127`
	RepetitionnumberR16     *PdschTimedomainresourceallocationR16RepetitionnumberR16
	K0V1710                 *utils.INTEGER                                             `lb:0,ub:128,ext`
	RepetitionnumberV1730   *PdschTimedomainresourceallocationR16RepetitionnumberV1730 `ext`
}

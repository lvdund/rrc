package ies

import "rrc/utils"

// PUSCH-TimeDomainResourceAllocation-r16 ::= SEQUENCE
// Extensible
type PuschTimedomainresourceallocationR16 struct {
	K2R16                  *utils.INTEGER       `lb:0,ub:32`
	PuschallocationlistR16 []PuschAllocationR16 `lb:1,ub:maxNrofMultiplePUSCHsR16`
}

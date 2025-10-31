package ies

import "rrc/utils"

// SRS-PosResourceSet-r16-resourceType-r16-aperiodic-r16 ::= SEQUENCE
// Extensible
type SrsPosresourcesetR16ResourcetypeR16AperiodicR16 struct {
	AperiodicsrsResourcetriggerlistR16 *[]utils.INTEGER `lb:1,ub:maxNrofSRSTriggerstates1`
}

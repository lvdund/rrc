package ies

import "rrc/utils"

// SRS-ResourceSet-resourceType-aperiodic ::= SEQUENCE
// Extensible
type SrsResourcesetResourcetypeAperiodic struct {
	AperiodicsrsResourcetrigger     utils.INTEGER `lb:0,ub:maxNrofSRSTriggerstates1`
	CsiRs                           *NzpCsiRsResourceid
	Slotoffset                      *utils.INTEGER   `lb:0,ub:32`
	AperiodicsrsResourcetriggerlist *[]utils.INTEGER `lb:1,ub:maxNrofSRSTriggerstates2`
}

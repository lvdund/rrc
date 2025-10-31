package ies

import "rrc/utils"

// DMRS-UplinkConfig ::= SEQUENCE
// Extensible
type DmrsUplinkconfig struct {
	DmrsType                   *DmrsUplinkconfigDmrsType
	DmrsAdditionalposition     *DmrsUplinkconfigDmrsAdditionalposition
	Phasetrackingrs            *utils.Setuprelease[PtrsUplinkconfig]
	Maxlength                  *DmrsUplinkconfigMaxlength
	Transformprecodingdisabled *DmrsUplinkconfigTransformprecodingdisabled `ext`
	Transformprecodingenabled  *DmrsUplinkconfigTransformprecodingenabled  `ext`
}

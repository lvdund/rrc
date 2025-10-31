package ies

import "rrc/utils"

// DMRS-DownlinkConfig ::= SEQUENCE
// Extensible
type DmrsDownlinkconfig struct {
	DmrsType               *DmrsDownlinkconfigDmrsType
	DmrsAdditionalposition *DmrsDownlinkconfigDmrsAdditionalposition
	Maxlength              *DmrsDownlinkconfigMaxlength
	Scramblingid0          *utils.INTEGER `lb:0,ub:65535`
	Scramblingid1          *utils.INTEGER `lb:0,ub:65535`
	Phasetrackingrs        *utils.Setuprelease[PtrsDownlinkconfig]
	DmrsDownlinkR16        *DmrsDownlinkconfigDmrsDownlinkR16 `ext`
}

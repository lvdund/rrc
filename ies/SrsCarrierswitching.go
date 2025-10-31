package ies

import "rrc/utils"

// SRS-CarrierSwitching ::= SEQUENCE
// Extensible
type SrsCarrierswitching struct {
	SrsSwitchfromservcellindex *utils.INTEGER `lb:0,ub:31`
	SrsSwitchfromcarrier       SrsCarrierswitchingSrsSwitchfromcarrier
	SrsTpcPdcchGroup           *SrsCarrierswitchingSrsTpcPdcchGroup
	Monitoringcells            *[]Servcellindex `lb:1,ub:maxNrofServingCells`
}

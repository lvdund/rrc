package ies

import "rrc/utils"

// SRS-Resources ::= SEQUENCE
type SrsResources struct {
	MaxnumberaperiodicsrsPerbwp             SrsResourcesMaxnumberaperiodicsrsPerbwp
	MaxnumberaperiodicsrsPerbwpPerslot      utils.INTEGER `lb:0,ub:6`
	MaxnumberperiodicsrsPerbwp              SrsResourcesMaxnumberperiodicsrsPerbwp
	MaxnumberperiodicsrsPerbwpPerslot       utils.INTEGER `lb:0,ub:6`
	MaxnumbersemipersistentsrsPerbwp        SrsResourcesMaxnumbersemipersistentsrsPerbwp
	MaxnumbersemipersistentsrsPerbwpPerslot utils.INTEGER `lb:0,ub:6`
	MaxnumbersrsPortsPerresource            SrsResourcesMaxnumbersrsPortsPerresource
}

package ies

import "rrc/utils"

// MeasObjectEUTRA ::= SEQUENCE
// Extensible
type Measobjecteutra struct {
	Carrierfreq                     ArfcnValueeutra
	Allowedmeasbandwidth            EutraAllowedmeasbandwidth
	Cellstoremovelisteutran         *EutraCellindexlist
	Cellstoaddmodlisteutran         *[]EutraCell `lb:1,ub:maxCellMeasEUTRA`
	Excludedcellstoremovelisteutran *EutraCellindexlist
	Excludedcellstoaddmodlisteutran *[]EutraExcludedcell `lb:1,ub:maxCellMeasEUTRA`
	EutraPresenceantennaport1       EutraPresenceantennaport1
	EutraQOffsetrange               *EutraQOffsetrange
	WidebandrsrqMeas                utils.BOOLEAN
	AssociatedmeasgapR17            *MeasgapidR17 `ext`
}

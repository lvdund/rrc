package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-v1610 ::= SEQUENCE
type NeighcellsiAcquisitionparametersV1610 struct {
	EutraSiAcquisitionforhoEndcR16 *utils.ENUMERATED
	NrAutonomousgapsEndcFr1R16     *utils.ENUMERATED
	NrAutonomousgapsEndcFr2R16     *utils.ENUMERATED
	NrAutonomousgapsFr1R16         *utils.ENUMERATED
	NrAutonomousgapsFr2R16         *utils.ENUMERATED
}

package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-v1610-nr-AutonomousGaps-FR1-r16 ::= ENUMERATED
type NeighcellsiAcquisitionparametersV1610NrAutonomousgapsFr1R16 struct {
	Value utils.ENUMERATED
}

const (
	NeighcellsiAcquisitionparametersV1610NrAutonomousgapsFr1R16EnumeratedNothing = iota
	NeighcellsiAcquisitionparametersV1610NrAutonomousgapsFr1R16EnumeratedSupported
)

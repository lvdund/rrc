package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-v1610-nr-AutonomousGaps-FR2-r16 ::= ENUMERATED
type NeighcellsiAcquisitionparametersV1610NrAutonomousgapsFr2R16 struct {
	Value utils.ENUMERATED
}

const (
	NeighcellsiAcquisitionparametersV1610NrAutonomousgapsFr2R16EnumeratedNothing = iota
	NeighcellsiAcquisitionparametersV1610NrAutonomousgapsFr2R16EnumeratedSupported
)

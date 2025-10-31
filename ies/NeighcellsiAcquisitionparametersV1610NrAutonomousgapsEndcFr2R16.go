package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-v1610-nr-AutonomousGaps-ENDC-FR2-r16 ::= ENUMERATED
type NeighcellsiAcquisitionparametersV1610NrAutonomousgapsEndcFr2R16 struct {
	Value utils.ENUMERATED
}

const (
	NeighcellsiAcquisitionparametersV1610NrAutonomousgapsEndcFr2R16EnumeratedNothing = iota
	NeighcellsiAcquisitionparametersV1610NrAutonomousgapsEndcFr2R16EnumeratedSupported
)

package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-r9-interFreqSI-AcquisitionForHO-r9 ::= ENUMERATED
type NeighcellsiAcquisitionparametersR9InterfreqsiAcquisitionforhoR9 struct {
	Value utils.ENUMERATED
}

const (
	NeighcellsiAcquisitionparametersR9InterfreqsiAcquisitionforhoR9EnumeratedNothing = iota
	NeighcellsiAcquisitionparametersR9InterfreqsiAcquisitionforhoR9EnumeratedSupported
)

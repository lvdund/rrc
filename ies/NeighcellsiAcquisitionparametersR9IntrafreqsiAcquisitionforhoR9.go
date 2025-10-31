package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-r9-intraFreqSI-AcquisitionForHO-r9 ::= ENUMERATED
type NeighcellsiAcquisitionparametersR9IntrafreqsiAcquisitionforhoR9 struct {
	Value utils.ENUMERATED
}

const (
	NeighcellsiAcquisitionparametersR9IntrafreqsiAcquisitionforhoR9EnumeratedNothing = iota
	NeighcellsiAcquisitionparametersR9IntrafreqsiAcquisitionforhoR9EnumeratedSupported
)

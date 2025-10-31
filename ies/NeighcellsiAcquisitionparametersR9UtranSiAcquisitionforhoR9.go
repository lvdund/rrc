package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-r9-utran-SI-AcquisitionForHO-r9 ::= ENUMERATED
type NeighcellsiAcquisitionparametersR9UtranSiAcquisitionforhoR9 struct {
	Value utils.ENUMERATED
}

const (
	NeighcellsiAcquisitionparametersR9UtranSiAcquisitionforhoR9EnumeratedNothing = iota
	NeighcellsiAcquisitionparametersR9UtranSiAcquisitionforhoR9EnumeratedSupported
)

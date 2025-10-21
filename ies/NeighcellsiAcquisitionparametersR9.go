package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-r9 ::= SEQUENCE
type NeighcellsiAcquisitionparametersR9 struct {
	IntrafreqsiAcquisitionforhoR9 *utils.ENUMERATED
	InterfreqsiAcquisitionforhoR9 *utils.ENUMERATED
	UtranSiAcquisitionforhoR9     *utils.ENUMERATED
}

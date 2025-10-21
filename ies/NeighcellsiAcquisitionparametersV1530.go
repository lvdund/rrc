package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-v1530 ::= SEQUENCE
type NeighcellsiAcquisitionparametersV1530 struct {
	ReportcgiNrEnDcR15   *utils.ENUMERATED
	ReportcgiNrNoenDcR15 *utils.ENUMERATED
}

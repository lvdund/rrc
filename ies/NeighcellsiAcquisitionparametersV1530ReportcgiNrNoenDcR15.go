package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-v1530-reportCGI-NR-NoEN-DC-r15 ::= ENUMERATED
type NeighcellsiAcquisitionparametersV1530ReportcgiNrNoenDcR15 struct {
	Value utils.ENUMERATED
}

const (
	NeighcellsiAcquisitionparametersV1530ReportcgiNrNoenDcR15EnumeratedNothing = iota
	NeighcellsiAcquisitionparametersV1530ReportcgiNrNoenDcR15EnumeratedSupported
)

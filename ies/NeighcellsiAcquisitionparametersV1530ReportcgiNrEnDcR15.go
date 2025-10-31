package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-v1530-reportCGI-NR-EN-DC-r15 ::= ENUMERATED
type NeighcellsiAcquisitionparametersV1530ReportcgiNrEnDcR15 struct {
	Value utils.ENUMERATED
}

const (
	NeighcellsiAcquisitionparametersV1530ReportcgiNrEnDcR15EnumeratedNothing = iota
	NeighcellsiAcquisitionparametersV1530ReportcgiNrEnDcR15EnumeratedSupported
)

package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-v1550-utra-GERAN-CGI-Reporting-ENDC-r15 ::= ENUMERATED
type NeighcellsiAcquisitionparametersV1550UtraGeranCgiReportingEndcR15 struct {
	Value utils.ENUMERATED
}

const (
	NeighcellsiAcquisitionparametersV1550UtraGeranCgiReportingEndcR15EnumeratedNothing = iota
	NeighcellsiAcquisitionparametersV1550UtraGeranCgiReportingEndcR15EnumeratedSupported
)

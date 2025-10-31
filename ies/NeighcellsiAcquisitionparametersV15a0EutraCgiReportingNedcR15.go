package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-v15a0-eutra-CGI-Reporting-NEDC-r15 ::= ENUMERATED
type NeighcellsiAcquisitionparametersV15a0EutraCgiReportingNedcR15 struct {
	Value utils.ENUMERATED
}

const (
	NeighcellsiAcquisitionparametersV15a0EutraCgiReportingNedcR15EnumeratedNothing = iota
	NeighcellsiAcquisitionparametersV15a0EutraCgiReportingNedcR15EnumeratedSupported
)

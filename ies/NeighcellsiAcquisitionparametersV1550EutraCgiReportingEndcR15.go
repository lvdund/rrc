package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-v1550-eutra-CGI-Reporting-ENDC-r15 ::= ENUMERATED
type NeighcellsiAcquisitionparametersV1550EutraCgiReportingEndcR15 struct {
	Value utils.ENUMERATED
}

const (
	NeighcellsiAcquisitionparametersV1550EutraCgiReportingEndcR15EnumeratedNothing = iota
	NeighcellsiAcquisitionparametersV1550EutraCgiReportingEndcR15EnumeratedSupported
)

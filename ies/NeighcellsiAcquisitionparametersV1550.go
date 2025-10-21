package ies

import "rrc/utils"

// NeighCellSI-AcquisitionParameters-v1550 ::= SEQUENCE
type NeighcellsiAcquisitionparametersV1550 struct {
	EutraCgiReportingEndcR15     *utils.ENUMERATED
	UtraGeranCgiReportingEndcR15 *utils.ENUMERATED
}

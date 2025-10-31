package ies

import "rrc/utils"

// MeasReportQuantityUTRA-FDD-r16 ::= SEQUENCE
type MeasreportquantityutraFddR16 struct {
	CpichRscp utils.BOOLEAN
	CpichEcn0 utils.BOOLEAN
}

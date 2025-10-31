package ies

import "rrc/utils"

// ReportSFTD-EUTRA ::= SEQUENCE
// Extensible
type ReportsftdEutra struct {
	ReportsftdMeas utils.BOOLEAN
	Reportrsrp     utils.BOOLEAN
}

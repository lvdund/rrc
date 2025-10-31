package ies

import "rrc/utils"

// ReportQuantityNR-r15 ::= SEQUENCE
type ReportquantitynrR15 struct {
	SsRsrp utils.BOOLEAN
	SsRsrq utils.BOOLEAN
	SsSinr utils.BOOLEAN
}

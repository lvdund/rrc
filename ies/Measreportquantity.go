package ies

import "rrc/utils"

// MeasReportQuantity ::= SEQUENCE
type Measreportquantity struct {
	Rsrp utils.BOOLEAN
	Rsrq utils.BOOLEAN
	Sinr utils.BOOLEAN
}

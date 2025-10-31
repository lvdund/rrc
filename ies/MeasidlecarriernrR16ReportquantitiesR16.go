package ies

import "rrc/utils"

// MeasIdleCarrierNR-r16-reportQuantities-r16 ::= ENUMERATED
type MeasidlecarriernrR16ReportquantitiesR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasidlecarriernrR16ReportquantitiesR16EnumeratedNothing = iota
	MeasidlecarriernrR16ReportquantitiesR16EnumeratedRsrp
	MeasidlecarriernrR16ReportquantitiesR16EnumeratedRsrq
	MeasidlecarriernrR16ReportquantitiesR16EnumeratedBoth
)

package ies

import "rrc/utils"

// MeasIdleCarrierNR-r16-reportQuantitiesNR-r16 ::= ENUMERATED
type MeasidlecarriernrR16ReportquantitiesnrR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasidlecarriernrR16ReportquantitiesnrR16EnumeratedNothing = iota
	MeasidlecarriernrR16ReportquantitiesnrR16EnumeratedRsrp
	MeasidlecarriernrR16ReportquantitiesnrR16EnumeratedRsrq
	MeasidlecarriernrR16ReportquantitiesnrR16EnumeratedBoth
)

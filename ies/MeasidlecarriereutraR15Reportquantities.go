package ies

import "rrc/utils"

// MeasIdleCarrierEUTRA-r15-reportQuantities ::= ENUMERATED
type MeasidlecarriereutraR15Reportquantities struct {
	Value utils.ENUMERATED
}

const (
	MeasidlecarriereutraR15ReportquantitiesEnumeratedNothing = iota
	MeasidlecarriereutraR15ReportquantitiesEnumeratedRsrp
	MeasidlecarriereutraR15ReportquantitiesEnumeratedRsrq
	MeasidlecarriereutraR15ReportquantitiesEnumeratedBoth
)

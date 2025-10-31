package ies

import "rrc/utils"

// MeasIdleCarrierEUTRA-r16-reportQuantitiesEUTRA-r16 ::= ENUMERATED
type MeasidlecarriereutraR16ReportquantitieseutraR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasidlecarriereutraR16ReportquantitieseutraR16EnumeratedNothing = iota
	MeasidlecarriereutraR16ReportquantitieseutraR16EnumeratedRsrp
	MeasidlecarriereutraR16ReportquantitieseutraR16EnumeratedRsrq
	MeasidlecarriereutraR16ReportquantitieseutraR16EnumeratedBoth
)

package ies

import "rrc/utils"

// MeasReportQuantityCLI-r16 ::= ENUMERATED
type MeasreportquantitycliR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasreportquantitycliR16EnumeratedNothing = iota
	MeasreportquantitycliR16EnumeratedSrs_Rsrp
	MeasreportquantitycliR16EnumeratedCli_Rssi
)

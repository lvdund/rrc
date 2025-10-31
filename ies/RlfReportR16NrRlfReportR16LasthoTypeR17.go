package ies

import "rrc/utils"

// RLF-Report-r16-nr-RLF-Report-r16-lastHO-Type-r17 ::= ENUMERATED
type RlfReportR16NrRlfReportR16LasthoTypeR17 struct {
	Value utils.ENUMERATED
}

const (
	RlfReportR16NrRlfReportR16LasthoTypeR17EnumeratedNothing = iota
	RlfReportR16NrRlfReportR16LasthoTypeR17EnumeratedCho
	RlfReportR16NrRlfReportR16LasthoTypeR17EnumeratedDaps
	RlfReportR16NrRlfReportR16LasthoTypeR17EnumeratedSpare2
	RlfReportR16NrRlfReportR16LasthoTypeR17EnumeratedSpare1
)

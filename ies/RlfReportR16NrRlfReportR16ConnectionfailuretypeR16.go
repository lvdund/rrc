package ies

import "rrc/utils"

// RLF-Report-r16-nr-RLF-Report-r16-connectionFailureType-r16 ::= ENUMERATED
type RlfReportR16NrRlfReportR16ConnectionfailuretypeR16 struct {
	Value utils.ENUMERATED
}

const (
	RlfReportR16NrRlfReportR16ConnectionfailuretypeR16EnumeratedNothing = iota
	RlfReportR16NrRlfReportR16ConnectionfailuretypeR16EnumeratedRlf
	RlfReportR16NrRlfReportR16ConnectionfailuretypeR16EnumeratedHof
)

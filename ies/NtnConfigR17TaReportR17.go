package ies

import "rrc/utils"

// NTN-Config-r17-ta-Report-r17 ::= ENUMERATED
type NtnConfigR17TaReportR17 struct {
	Value utils.ENUMERATED
}

const (
	NtnConfigR17TaReportR17EnumeratedNothing = iota
	NtnConfigR17TaReportR17EnumeratedEnabled
)

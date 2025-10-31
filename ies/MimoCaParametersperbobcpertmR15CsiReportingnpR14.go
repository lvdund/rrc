package ies

import "rrc/utils"

// MIMO-CA-ParametersPerBoBCPerTM-r15-csi-ReportingNP-r14 ::= ENUMERATED
type MimoCaParametersperbobcpertmR15CsiReportingnpR14 struct {
	Value utils.ENUMERATED
}

const (
	MimoCaParametersperbobcpertmR15CsiReportingnpR14EnumeratedNothing = iota
	MimoCaParametersperbobcpertmR15CsiReportingnpR14EnumeratedDifferent
)

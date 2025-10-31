package ies

import "rrc/utils"

// MIMO-CA-ParametersPerBoBCPerTM-v1430-csi-ReportingNP-r14 ::= ENUMERATED
type MimoCaParametersperbobcpertmV1430CsiReportingnpR14 struct {
	Value utils.ENUMERATED
}

const (
	MimoCaParametersperbobcpertmV1430CsiReportingnpR14EnumeratedNothing = iota
	MimoCaParametersperbobcpertmV1430CsiReportingnpR14EnumeratedDifferent
)

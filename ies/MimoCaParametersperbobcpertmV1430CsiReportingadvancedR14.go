package ies

import "rrc/utils"

// MIMO-CA-ParametersPerBoBCPerTM-v1430-csi-ReportingAdvanced-r14 ::= ENUMERATED
type MimoCaParametersperbobcpertmV1430CsiReportingadvancedR14 struct {
	Value utils.ENUMERATED
}

const (
	MimoCaParametersperbobcpertmV1430CsiReportingadvancedR14EnumeratedNothing = iota
	MimoCaParametersperbobcpertmV1430CsiReportingadvancedR14EnumeratedDifferent
)

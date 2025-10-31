package ies

import "rrc/utils"

// MIMO-CA-ParametersPerBoBCPerTM-r15-csi-ReportingAdvanced-r14 ::= ENUMERATED
type MimoCaParametersperbobcpertmR15CsiReportingadvancedR14 struct {
	Value utils.ENUMERATED
}

const (
	MimoCaParametersperbobcpertmR15CsiReportingadvancedR14EnumeratedNothing = iota
	MimoCaParametersperbobcpertmR15CsiReportingadvancedR14EnumeratedDifferent
)

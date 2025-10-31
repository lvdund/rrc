package ies

import "rrc/utils"

// MIMO-UE-ParametersPerTM-v1430-csi-ReportingAdvanced-r14 ::= ENUMERATED
type MimoUeParameterspertmV1430CsiReportingadvancedR14 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeParameterspertmV1430CsiReportingadvancedR14EnumeratedNothing = iota
	MimoUeParameterspertmV1430CsiReportingadvancedR14EnumeratedSupported
)

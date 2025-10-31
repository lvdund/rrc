package ies

import "rrc/utils"

// MIMO-UE-ParametersPerTM-v1430-csi-ReportingNP-r14 ::= ENUMERATED
type MimoUeParameterspertmV1430CsiReportingnpR14 struct {
	Value utils.ENUMERATED
}

const (
	MimoUeParameterspertmV1430CsiReportingnpR14EnumeratedNothing = iota
	MimoUeParameterspertmV1430CsiReportingnpR14EnumeratedSupported
)

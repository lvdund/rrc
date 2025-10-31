package ies

import "rrc/utils"

// MIMO-ParametersPerBand-groupSINR-reporting-r16 ::= ENUMERATED
type MimoParametersperbandGroupsinrReportingR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandGroupsinrReportingR16EnumeratedNothing = iota
	MimoParametersperbandGroupsinrReportingR16EnumeratedSupported
)

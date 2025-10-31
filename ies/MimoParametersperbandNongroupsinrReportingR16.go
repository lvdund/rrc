package ies

import "rrc/utils"

// MIMO-ParametersPerBand-nonGroupSINR-reporting-r16 ::= ENUMERATED
type MimoParametersperbandNongroupsinrReportingR16 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandNongroupsinrReportingR16EnumeratedNothing = iota
	MimoParametersperbandNongroupsinrReportingR16EnumeratedN1
	MimoParametersperbandNongroupsinrReportingR16EnumeratedN2
	MimoParametersperbandNongroupsinrReportingR16EnumeratedN4
)

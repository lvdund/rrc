package ies

import "rrc/utils"

// MAC-ParametersCommon-tdd-MPE-P-MPR-Reporting-r16 ::= ENUMERATED
type MacParameterscommonTddMpePMprReportingR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonTddMpePMprReportingR16EnumeratedNothing = iota
	MacParameterscommonTddMpePMprReportingR16EnumeratedSupported
)

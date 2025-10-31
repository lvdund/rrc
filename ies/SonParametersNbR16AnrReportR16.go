package ies

import "rrc/utils"

// SON-Parameters-NB-r16-anr-Report-r16 ::= ENUMERATED
type SonParametersNbR16AnrReportR16 struct {
	Value utils.ENUMERATED
}

const (
	SonParametersNbR16AnrReportR16EnumeratedNothing = iota
	SonParametersNbR16AnrReportR16EnumeratedSupported
)

package ies

import "rrc/utils"

// SON-Parameters-NB-r16-rach-Report-r16 ::= ENUMERATED
type SonParametersNbR16RachReportR16 struct {
	Value utils.ENUMERATED
}

const (
	SonParametersNbR16RachReportR16EnumeratedNothing = iota
	SonParametersNbR16RachReportR16EnumeratedSupported
)

package ies

import "rrc/utils"

// SON-Parameters-r16-rach-Report-r16 ::= ENUMERATED
type SonParametersR16RachReportR16 struct {
	Value utils.ENUMERATED
}

const (
	SonParametersR16RachReportR16EnumeratedNothing = iota
	SonParametersR16RachReportR16EnumeratedSupported
)

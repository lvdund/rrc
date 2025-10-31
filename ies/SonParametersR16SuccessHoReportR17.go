package ies

import "rrc/utils"

// SON-Parameters-r16-success-HO-Report-r17 ::= ENUMERATED
type SonParametersR16SuccessHoReportR17 struct {
	Value utils.ENUMERATED
}

const (
	SonParametersR16SuccessHoReportR17EnumeratedNothing = iota
	SonParametersR16SuccessHoReportR17EnumeratedSupported
)

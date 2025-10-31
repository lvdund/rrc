package ies

import "rrc/utils"

// SON-Parameters-r16-twoStepRACH-Report-r17 ::= ENUMERATED
type SonParametersR16TwosteprachReportR17 struct {
	Value utils.ENUMERATED
}

const (
	SonParametersR16TwosteprachReportR17EnumeratedNothing = iota
	SonParametersR16TwosteprachReportR17EnumeratedSupported
)

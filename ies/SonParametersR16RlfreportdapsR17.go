package ies

import "rrc/utils"

// SON-Parameters-r16-rlfReportDAPS-r17 ::= ENUMERATED
type SonParametersR16RlfreportdapsR17 struct {
	Value utils.ENUMERATED
}

const (
	SonParametersR16RlfreportdapsR17EnumeratedNothing = iota
	SonParametersR16RlfreportdapsR17EnumeratedSupported
)

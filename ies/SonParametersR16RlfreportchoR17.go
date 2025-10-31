package ies

import "rrc/utils"

// SON-Parameters-r16-rlfReportCHO-r17 ::= ENUMERATED
type SonParametersR16RlfreportchoR17 struct {
	Value utils.ENUMERATED
}

const (
	SonParametersR16RlfreportchoR17EnumeratedNothing = iota
	SonParametersR16RlfreportchoR17EnumeratedSupported
)

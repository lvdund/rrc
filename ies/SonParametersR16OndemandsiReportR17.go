package ies

import "rrc/utils"

// SON-Parameters-r16-onDemandSI-Report-r17 ::= ENUMERATED
type SonParametersR16OndemandsiReportR17 struct {
	Value utils.ENUMERATED
}

const (
	SonParametersR16OndemandsiReportR17EnumeratedNothing = iota
	SonParametersR16OndemandsiReportR17EnumeratedSupported
)

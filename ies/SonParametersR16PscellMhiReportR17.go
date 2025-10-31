package ies

import "rrc/utils"

// SON-Parameters-r16-pscell-MHI-Report-r17 ::= ENUMERATED
type SonParametersR16PscellMhiReportR17 struct {
	Value utils.ENUMERATED
}

const (
	SonParametersR16PscellMhiReportR17EnumeratedNothing = iota
	SonParametersR16PscellMhiReportR17EnumeratedSupported
)

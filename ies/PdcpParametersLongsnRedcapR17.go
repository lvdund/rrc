package ies

import "rrc/utils"

// PDCP-Parameters-longSN-RedCap-r17 ::= ENUMERATED
type PdcpParametersLongsnRedcapR17 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersLongsnRedcapR17EnumeratedNothing = iota
	PdcpParametersLongsnRedcapR17EnumeratedSupported
)

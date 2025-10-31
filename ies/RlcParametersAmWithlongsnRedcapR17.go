package ies

import "rrc/utils"

// RLC-Parameters-am-WithLongSN-RedCap-r17 ::= ENUMERATED
type RlcParametersAmWithlongsnRedcapR17 struct {
	Value utils.ENUMERATED
}

const (
	RlcParametersAmWithlongsnRedcapR17EnumeratedNothing = iota
	RlcParametersAmWithlongsnRedcapR17EnumeratedSupported
)

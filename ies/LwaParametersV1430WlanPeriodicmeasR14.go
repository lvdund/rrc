package ies

import "rrc/utils"

// LWA-Parameters-v1430-wlan-PeriodicMeas-r14 ::= ENUMERATED
type LwaParametersV1430WlanPeriodicmeasR14 struct {
	Value utils.ENUMERATED
}

const (
	LwaParametersV1430WlanPeriodicmeasR14EnumeratedNothing = iota
	LwaParametersV1430WlanPeriodicmeasR14EnumeratedSupported
)

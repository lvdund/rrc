package ies

import "rrc/utils"

// SL-Parameters-v1310-discPeriodicSLSS-r13 ::= ENUMERATED
type SlParametersV1310DiscperiodicslssR13 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1310DiscperiodicslssR13EnumeratedNothing = iota
	SlParametersV1310DiscperiodicslssR13EnumeratedSupported
)

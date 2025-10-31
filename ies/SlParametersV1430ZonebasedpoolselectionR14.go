package ies

import "rrc/utils"

// SL-Parameters-v1430-zoneBasedPoolSelection-r14 ::= ENUMERATED
type SlParametersV1430ZonebasedpoolselectionR14 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1430ZonebasedpoolselectionR14EnumeratedNothing = iota
	SlParametersV1430ZonebasedpoolselectionR14EnumeratedSupported
)

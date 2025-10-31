package ies

import "rrc/utils"

// Other-Parameters-v1530-flightPathPlan-r15 ::= ENUMERATED
type OtherParametersV1530FlightpathplanR15 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1530FlightpathplanR15EnumeratedNothing = iota
	OtherParametersV1530FlightpathplanR15EnumeratedSupported
)

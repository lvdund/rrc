package ies

import "rrc/utils"

// MeasParameters-v1310-ul-PDCP-Delay-r13 ::= ENUMERATED
type MeasparametersV1310UlPdcpDelayR13 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1310UlPdcpDelayR13EnumeratedNothing = iota
	MeasparametersV1310UlPdcpDelayR13EnumeratedSupported
)

package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-v1430-ce-PUSCH-MaxBandwidth-r14 ::= ENUMERATED
type PuschConfigdedicatedV1430CePuschMaxbandwidthR14 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigdedicatedV1430CePuschMaxbandwidthR14EnumeratedNothing = iota
	PuschConfigdedicatedV1430CePuschMaxbandwidthR14EnumeratedBw5
)

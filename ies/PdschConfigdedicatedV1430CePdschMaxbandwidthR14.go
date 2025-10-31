package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1430-ce-PDSCH-MaxBandwidth-r14 ::= ENUMERATED
type PdschConfigdedicatedV1430CePdschMaxbandwidthR14 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigdedicatedV1430CePdschMaxbandwidthR14EnumeratedNothing = iota
	PdschConfigdedicatedV1430CePdschMaxbandwidthR14EnumeratedBw5
	PdschConfigdedicatedV1430CePdschMaxbandwidthR14EnumeratedBw20
)

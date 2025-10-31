package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-v1430-ce-PUSCH-NB-MaxTBS-r14 ::= ENUMERATED
type PuschConfigdedicatedV1430CePuschNbMaxtbsR14 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigdedicatedV1430CePuschNbMaxtbsR14EnumeratedNothing = iota
	PuschConfigdedicatedV1430CePuschNbMaxtbsR14EnumeratedOn
)

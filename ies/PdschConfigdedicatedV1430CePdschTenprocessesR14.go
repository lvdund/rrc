package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1430-ce-PDSCH-TenProcesses-r14 ::= ENUMERATED
type PdschConfigdedicatedV1430CePdschTenprocessesR14 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigdedicatedV1430CePdschTenprocessesR14EnumeratedNothing = iota
	PdschConfigdedicatedV1430CePdschTenprocessesR14EnumeratedOn
)

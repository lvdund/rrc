package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1430-ce-HARQ-AckBundling-r14 ::= ENUMERATED
type PdschConfigdedicatedV1430CeHarqAckbundlingR14 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigdedicatedV1430CeHarqAckbundlingR14EnumeratedNothing = iota
	PdschConfigdedicatedV1430CeHarqAckbundlingR14EnumeratedOn
)

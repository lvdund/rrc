package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1530-ce-PDSCH-64QAM-Config-r15 ::= ENUMERATED
type PdschConfigdedicatedV1530CePdsch64qamConfigR15 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigdedicatedV1530CePdsch64qamConfigR15EnumeratedNothing = iota
	PdschConfigdedicatedV1530CePdsch64qamConfigR15EnumeratedOn
)

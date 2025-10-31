package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1430-tbsIndexAlt2-r14 ::= ENUMERATED
type PdschConfigdedicatedV1430Tbsindexalt2R14 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigdedicatedV1430Tbsindexalt2R14EnumeratedNothing = iota
	PdschConfigdedicatedV1430Tbsindexalt2R14EnumeratedB33
)

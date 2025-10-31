package ies

import "rrc/utils"

// PDSCH-ConfigDedicatedSCell-v1430-tbsIndexAlt2-r14 ::= ENUMERATED
type PdschConfigdedicatedscellV1430Tbsindexalt2R14 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigdedicatedscellV1430Tbsindexalt2R14EnumeratedNothing = iota
	PdschConfigdedicatedscellV1430Tbsindexalt2R14EnumeratedB33
)

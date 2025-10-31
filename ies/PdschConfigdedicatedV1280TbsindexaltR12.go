package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1280-tbsIndexAlt-r12 ::= ENUMERATED
type PdschConfigdedicatedV1280TbsindexaltR12 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigdedicatedV1280TbsindexaltR12EnumeratedNothing = iota
	PdschConfigdedicatedV1280TbsindexaltR12EnumeratedA26
	PdschConfigdedicatedV1280TbsindexaltR12EnumeratedA33
)

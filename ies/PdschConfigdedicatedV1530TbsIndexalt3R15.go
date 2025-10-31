package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1530-tbs-IndexAlt3-r15 ::= ENUMERATED
type PdschConfigdedicatedV1530TbsIndexalt3R15 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigdedicatedV1530TbsIndexalt3R15EnumeratedNothing = iota
	PdschConfigdedicatedV1530TbsIndexalt3R15EnumeratedA37
)

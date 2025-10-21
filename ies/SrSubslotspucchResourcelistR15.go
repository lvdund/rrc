package ies

import "rrc/utils"

// SR-SubslotSPUCCH-ResourceList-r15 ::= SEQUENCE OF INTEGER
// SIZE (1..4)
type SrSubslotspucchResourcelistR15 struct {
	Value utils.Sequence[Integer]
}

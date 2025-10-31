package ies

import "rrc/utils"

// SR-SubslotSPUCCH-ResourceList-r15 ::= SEQUENCE OF utils.INTEGER // SIZE (1..4)
type SrSubslotspucchResourcelistR15 struct {
	Value []utils.INTEGER `lb:1,ub:4`
}

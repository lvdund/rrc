package ies

import "rrc/utils"

// SchedulingInfo-BR-r13 ::= SEQUENCE
type SchedulinginfoBrR13 struct {
	SiNarrowbandR13 utils.INTEGER `lb:0,ub:maxAvailNarrowBandsR13`
	SiTbsR13        SchedulinginfoBrR13SiTbsR13
}

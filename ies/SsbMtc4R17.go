package ies

import "rrc/utils"

// SSB-MTC4-r17 ::= SEQUENCE
type SsbMtc4R17 struct {
	PciListR17 *[]Physcellid `lb:1,ub:maxNrofPCIsPerSMTC`
	OffsetR17  utils.INTEGER `lb:0,ub:159`
}

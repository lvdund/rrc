package ies

import "rrc/utils"

// PUCCH-ResourceSet ::= SEQUENCE
type PucchResourceset struct {
	PucchResourcesetid PucchResourcesetid
	Resourcelist       []PucchResourceid `lb:1,ub:maxNrofPUCCHResourcesperset`
	Maxpayloadsize     *utils.INTEGER    `lb:0,ub:256`
}

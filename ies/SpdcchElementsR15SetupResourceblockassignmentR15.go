package ies

import "rrc/utils"

// SPDCCH-Elements-r15-setup-resourceBlockAssignment-r15 ::= SEQUENCE
type SpdcchElementsR15SetupResourceblockassignmentR15 struct {
	NumberrbInfreqDomainR15    utils.INTEGER   `lb:0,ub:100`
	ResourceblockassignmentR15 utils.BITSTRING `lb:98,ub:98`
}

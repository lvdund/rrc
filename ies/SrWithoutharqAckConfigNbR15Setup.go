package ies

import "rrc/utils"

// SR-WithoutHARQ-ACK-Config-NB-r15-setup ::= SEQUENCE
type SrWithoutharqAckConfigNbR15Setup struct {
	SrProhibittimerR15  *utils.INTEGER `lb:0,ub:7`
	SrNprachResourceR15 *SrNprachResourceNbR15
}

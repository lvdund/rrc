package ies

import "rrc/utils"

// T-PollRetransmit-NB-r13 ::= ENUMERATED
type TPollretransmitNbR13 struct {
	Value utils.ENUMERATED
}

const (
	TPollretransmitNbR13Ms250         = 0
	TPollretransmitNbR13Ms500         = 1
	TPollretransmitNbR13Ms1000        = 2
	TPollretransmitNbR13Ms2000        = 3
	TPollretransmitNbR13Ms3000        = 4
	TPollretransmitNbR13Ms4000        = 5
	TPollretransmitNbR13Ms6000        = 6
	TPollretransmitNbR13Ms10000       = 7
	TPollretransmitNbR13Ms15000       = 8
	TPollretransmitNbR13Ms25000       = 9
	TPollretransmitNbR13Ms40000       = 10
	TPollretransmitNbR13Ms60000       = 11
	TPollretransmitNbR13Ms90000       = 12
	TPollretransmitNbR13Ms120000      = 13
	TPollretransmitNbR13Ms180000      = 14
	TPollretransmitNbR13Ms300000V1530 = 15
)

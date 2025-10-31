package ies

import "rrc/utils"

// BWP ::= SEQUENCE
type Bwp struct {
	Locationandbandwidth utils.INTEGER `lb:0,ub:37949`
	Subcarrierspacing    Subcarrierspacing
	Cyclicprefix         *BwpCyclicprefix
}

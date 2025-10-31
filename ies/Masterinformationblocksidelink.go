package ies

import "rrc/utils"

// MasterInformationBlockSidelink ::= SEQUENCE
type Masterinformationblocksidelink struct {
	SlTddConfigR16       utils.BITSTRING `lb:12,ub:12`
	IncoverageR16        utils.BOOLEAN
	DirectframenumberR16 utils.BITSTRING `lb:10,ub:10`
	SlotindexR16         utils.BITSTRING `lb:7,ub:7`
	ReservedbitsR16      utils.BITSTRING `lb:2,ub:2`
}

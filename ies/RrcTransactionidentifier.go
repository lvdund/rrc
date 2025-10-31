package ies

import "rrc/utils"

// RRC-TransactionIdentifier ::= utils.INTEGER (0..3)
type RrcTransactionidentifier struct {
	Value utils.INTEGER `lb:0,ub:3`
}

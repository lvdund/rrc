package ies

import "rrc/utils"

// BandParameters-v1380 ::= SEQUENCE
type BandparametersV1380 struct {
	TxantennaswitchdlR13 *utils.INTEGER `lb:0,ub:32`
	TxantennaswitchulR13 *utils.INTEGER `lb:0,ub:32`
}

package ies

import "rrc/utils"

// LogMeasResultBT-r15 ::= SEQUENCE
// Extensible
type LogmeasresultbtR15 struct {
	BtAddrR15 utils.BITSTRING `lb:48,ub:48`
	RssiBtR15 *utils.INTEGER  `lb:0,ub:127`
}

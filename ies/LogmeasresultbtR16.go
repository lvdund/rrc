package ies

import "rrc/utils"

// LogMeasResultBT-r16 ::= SEQUENCE
// Extensible
type LogmeasresultbtR16 struct {
	BtAddrR16 utils.BITSTRING `lb:48,ub:48`
	RssiBtR16 *utils.INTEGER  `lb:0,ub:127`
}

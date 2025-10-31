package ies

import "rrc/utils"

// NR-TimeStamp-r17 ::= SEQUENCE
// Extensible
type NrTimestampR17 struct {
	NrSfnR17  utils.INTEGER `lb:0,ub:1023`
	NrSlotR17 NrTimestampR17NrSlotR17
}

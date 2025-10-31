package ies

import "rrc/utils"

// NR-TimeStamp-r17-nr-Slot-r17 ::= CHOICE
const (
	NrTimestampR17NrSlotR17ChoiceNothing = iota
	NrTimestampR17NrSlotR17ChoiceScs15R17
	NrTimestampR17NrSlotR17ChoiceScs30R17
	NrTimestampR17NrSlotR17ChoiceScs60R17
	NrTimestampR17NrSlotR17ChoiceScs120R17
)

type NrTimestampR17NrSlotR17 struct {
	Choice    uint64
	Scs15R17  *utils.INTEGER `lb:0,ub:9`
	Scs30R17  *utils.INTEGER `lb:0,ub:19`
	Scs60R17  *utils.INTEGER `lb:0,ub:39`
	Scs120R17 *utils.INTEGER `lb:0,ub:79`
}

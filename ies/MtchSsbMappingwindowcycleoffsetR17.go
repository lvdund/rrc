package ies

import "rrc/utils"

// MTCH-SSB-MappingWindowCycleOffset-r17 ::= CHOICE
const (
	MtchSsbMappingwindowcycleoffsetR17ChoiceNothing = iota
	MtchSsbMappingwindowcycleoffsetR17ChoiceMs10
	MtchSsbMappingwindowcycleoffsetR17ChoiceMs20
	MtchSsbMappingwindowcycleoffsetR17ChoiceMs32
	MtchSsbMappingwindowcycleoffsetR17ChoiceMs64
	MtchSsbMappingwindowcycleoffsetR17ChoiceMs128
	MtchSsbMappingwindowcycleoffsetR17ChoiceMs256
)

type MtchSsbMappingwindowcycleoffsetR17 struct {
	Choice uint64
	Ms10   *utils.INTEGER `lb:0,ub:9`
	Ms20   *utils.INTEGER `lb:0,ub:19`
	Ms32   *utils.INTEGER `lb:0,ub:31`
	Ms64   *utils.INTEGER `lb:0,ub:63`
	Ms128  *utils.INTEGER `lb:0,ub:127`
	Ms256  *utils.INTEGER `lb:0,ub:255`
}

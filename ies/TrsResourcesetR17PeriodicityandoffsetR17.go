package ies

import "rrc/utils"

// TRS-ResourceSet-r17-periodicityAndOffset-r17 ::= CHOICE
const (
	TrsResourcesetR17PeriodicityandoffsetR17ChoiceNothing = iota
	TrsResourcesetR17PeriodicityandoffsetR17ChoiceSlots10
	TrsResourcesetR17PeriodicityandoffsetR17ChoiceSlots20
	TrsResourcesetR17PeriodicityandoffsetR17ChoiceSlots40
	TrsResourcesetR17PeriodicityandoffsetR17ChoiceSlots80
)

type TrsResourcesetR17PeriodicityandoffsetR17 struct {
	Choice  uint64
	Slots10 *utils.INTEGER `lb:0,ub:9`
	Slots20 *utils.INTEGER `lb:0,ub:19`
	Slots40 *utils.INTEGER `lb:0,ub:39`
	Slots80 *utils.INTEGER `lb:0,ub:79`
}

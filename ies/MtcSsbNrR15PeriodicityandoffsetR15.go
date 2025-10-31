package ies

import "rrc/utils"

// MTC-SSB-NR-r15-periodicityAndOffset-r15 ::= CHOICE
const (
	MtcSsbNrR15PeriodicityandoffsetR15ChoiceNothing = iota
	MtcSsbNrR15PeriodicityandoffsetR15ChoiceSf5R15
	MtcSsbNrR15PeriodicityandoffsetR15ChoiceSf10R15
	MtcSsbNrR15PeriodicityandoffsetR15ChoiceSf20R15
	MtcSsbNrR15PeriodicityandoffsetR15ChoiceSf40R15
	MtcSsbNrR15PeriodicityandoffsetR15ChoiceSf80R15
	MtcSsbNrR15PeriodicityandoffsetR15ChoiceSf160R15
)

type MtcSsbNrR15PeriodicityandoffsetR15 struct {
	Choice   uint64
	Sf5R15   *utils.INTEGER `lb:0,ub:4`
	Sf10R15  *utils.INTEGER `lb:0,ub:9`
	Sf20R15  *utils.INTEGER `lb:0,ub:19`
	Sf40R15  *utils.INTEGER `lb:0,ub:39`
	Sf80R15  *utils.INTEGER `lb:0,ub:79`
	Sf160R15 *utils.INTEGER `lb:0,ub:159`
}

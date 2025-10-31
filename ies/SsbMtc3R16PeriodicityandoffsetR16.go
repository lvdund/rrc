package ies

import "rrc/utils"

// SSB-MTC3-r16-periodicityAndOffset-r16 ::= CHOICE
const (
	SsbMtc3R16PeriodicityandoffsetR16ChoiceNothing = iota
	SsbMtc3R16PeriodicityandoffsetR16ChoiceSf5R16
	SsbMtc3R16PeriodicityandoffsetR16ChoiceSf10R16
	SsbMtc3R16PeriodicityandoffsetR16ChoiceSf20R16
	SsbMtc3R16PeriodicityandoffsetR16ChoiceSf40R16
	SsbMtc3R16PeriodicityandoffsetR16ChoiceSf80R16
	SsbMtc3R16PeriodicityandoffsetR16ChoiceSf160R16
	SsbMtc3R16PeriodicityandoffsetR16ChoiceSf320R16
	SsbMtc3R16PeriodicityandoffsetR16ChoiceSf640R16
	SsbMtc3R16PeriodicityandoffsetR16ChoiceSf1280R16
)

type SsbMtc3R16PeriodicityandoffsetR16 struct {
	Choice    uint64
	Sf5R16    *utils.INTEGER `lb:0,ub:4`
	Sf10R16   *utils.INTEGER `lb:0,ub:9`
	Sf20R16   *utils.INTEGER `lb:0,ub:19`
	Sf40R16   *utils.INTEGER `lb:0,ub:39`
	Sf80R16   *utils.INTEGER `lb:0,ub:79`
	Sf160R16  *utils.INTEGER `lb:0,ub:159`
	Sf320R16  *utils.INTEGER `lb:0,ub:319`
	Sf640R16  *utils.INTEGER `lb:0,ub:639`
	Sf1280R16 *utils.INTEGER `lb:0,ub:1279`
}

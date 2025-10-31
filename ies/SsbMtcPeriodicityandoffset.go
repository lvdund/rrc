package ies

import "rrc/utils"

// SSB-MTC-periodicityAndOffset ::= CHOICE
const (
	SsbMtcPeriodicityandoffsetChoiceNothing = iota
	SsbMtcPeriodicityandoffsetChoiceSf5
	SsbMtcPeriodicityandoffsetChoiceSf10
	SsbMtcPeriodicityandoffsetChoiceSf20
	SsbMtcPeriodicityandoffsetChoiceSf40
	SsbMtcPeriodicityandoffsetChoiceSf80
	SsbMtcPeriodicityandoffsetChoiceSf160
)

type SsbMtcPeriodicityandoffset struct {
	Choice uint64
	Sf5    *utils.INTEGER `lb:0,ub:4`
	Sf10   *utils.INTEGER `lb:0,ub:9`
	Sf20   *utils.INTEGER `lb:0,ub:19`
	Sf40   *utils.INTEGER `lb:0,ub:39`
	Sf80   *utils.INTEGER `lb:0,ub:79`
	Sf160  *utils.INTEGER `lb:0,ub:159`
}

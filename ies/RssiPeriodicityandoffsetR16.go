package ies

import "rrc/utils"

// RSSI-PeriodicityAndOffset-r16 ::= CHOICE
// Extensible
const (
	RssiPeriodicityandoffsetR16ChoiceNothing = iota
	RssiPeriodicityandoffsetR16ChoiceSl10
	RssiPeriodicityandoffsetR16ChoiceSl20
	RssiPeriodicityandoffsetR16ChoiceSl40
	RssiPeriodicityandoffsetR16ChoiceSl80
	RssiPeriodicityandoffsetR16ChoiceSl160
	RssiPeriodicityandoffsetR16ChoiceSl320
	RssiPeriodicityandoffsetR16ChoiceS1640
)

type RssiPeriodicityandoffsetR16 struct {
	Choice uint64
	Sl10   *utils.INTEGER `lb:0,ub:9`
	Sl20   *utils.INTEGER `lb:0,ub:19`
	Sl40   *utils.INTEGER `lb:0,ub:39`
	Sl80   *utils.INTEGER `lb:0,ub:79`
	Sl160  *utils.INTEGER `lb:0,ub:159`
	Sl320  *utils.INTEGER `lb:0,ub:319`
	S1640  *utils.INTEGER `lb:0,ub:639`
}

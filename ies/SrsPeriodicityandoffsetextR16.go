package ies

import "rrc/utils"

// SRS-PeriodicityAndOffsetExt-r16 ::= CHOICE
const (
	SrsPeriodicityandoffsetextR16ChoiceNothing = iota
	SrsPeriodicityandoffsetextR16ChoiceSl128
	SrsPeriodicityandoffsetextR16ChoiceSl256
	SrsPeriodicityandoffsetextR16ChoiceSl512
	SrsPeriodicityandoffsetextR16ChoiceSl20480
)

type SrsPeriodicityandoffsetextR16 struct {
	Choice  uint64
	Sl128   *utils.INTEGER `lb:0,ub:127`
	Sl256   *utils.INTEGER `lb:0,ub:255`
	Sl512   *utils.INTEGER `lb:0,ub:511`
	Sl20480 *utils.INTEGER `lb:0,ub:20479`
}

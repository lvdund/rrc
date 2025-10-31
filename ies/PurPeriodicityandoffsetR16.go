package ies

import "rrc/utils"

// PUR-PeriodicityAndOffset-r16 ::= CHOICE
const (
	PurPeriodicityandoffsetR16ChoiceNothing = iota
	PurPeriodicityandoffsetR16ChoicePeriodicity8
	PurPeriodicityandoffsetR16ChoicePeriodicity16
	PurPeriodicityandoffsetR16ChoicePeriodicity32
	PurPeriodicityandoffsetR16ChoicePeriodicity64
	PurPeriodicityandoffsetR16ChoicePeriodicity128
	PurPeriodicityandoffsetR16ChoicePeriodicity256
	PurPeriodicityandoffsetR16ChoicePeriodicity512
	PurPeriodicityandoffsetR16ChoicePeriodicity1024
	PurPeriodicityandoffsetR16ChoicePeriodicity2048
	PurPeriodicityandoffsetR16ChoicePeriodicity4096
	PurPeriodicityandoffsetR16ChoicePeriodicity8192
)

type PurPeriodicityandoffsetR16 struct {
	Choice          uint64
	Periodicity8    *utils.INTEGER `lb:1,ub:7`
	Periodicity16   *utils.INTEGER `lb:1,ub:15`
	Periodicity32   *utils.INTEGER `lb:1,ub:31`
	Periodicity64   *utils.INTEGER `lb:1,ub:63`
	Periodicity128  *utils.INTEGER `lb:1,ub:127`
	Periodicity256  *utils.INTEGER `lb:1,ub:255`
	Periodicity512  *utils.INTEGER `lb:1,ub:511`
	Periodicity1024 *utils.INTEGER `lb:1,ub:1023`
	Periodicity2048 *utils.INTEGER `lb:1,ub:2047`
	Periodicity4096 *utils.INTEGER `lb:1,ub:4095`
	Periodicity8192 *utils.INTEGER `lb:1,ub:8191`
}

package ies

import "rrc/utils"

// PUR-PeriodicityAndOffset-NB-r16 ::= CHOICE
const (
	PurPeriodicityandoffsetNbR16ChoiceNothing = iota
	PurPeriodicityandoffsetNbR16ChoicePeriodicity8
	PurPeriodicityandoffsetNbR16ChoicePeriodicity16
	PurPeriodicityandoffsetNbR16ChoicePeriodicity32
	PurPeriodicityandoffsetNbR16ChoicePeriodicity64
	PurPeriodicityandoffsetNbR16ChoicePeriodicity128
	PurPeriodicityandoffsetNbR16ChoicePeriodicity256
	PurPeriodicityandoffsetNbR16ChoicePeriodicity512
	PurPeriodicityandoffsetNbR16ChoicePeriodicity1024
	PurPeriodicityandoffsetNbR16ChoicePeriodicity2048
	PurPeriodicityandoffsetNbR16ChoicePeriodicity4096
	PurPeriodicityandoffsetNbR16ChoicePeriodicity8192
)

type PurPeriodicityandoffsetNbR16 struct {
	Choice          uint64
	Periodicity8    *utils.INTEGER `lb:1,ub:7`
	Periodicity16   *utils.INTEGER `lb:1,ub:15`
	Periodicity32   *utils.INTEGER `lb:1,ub:31`
	Periodicity64   *utils.INTEGER `lb:1,ub:63`
	Periodicity128  *utils.INTEGER `lb:1,ub:127`
	Periodicity256  *utils.INTEGER `lb:1,ub:257`
	Periodicity512  *utils.INTEGER `lb:1,ub:511`
	Periodicity1024 *utils.INTEGER `lb:1,ub:1023`
	Periodicity2048 *utils.INTEGER `lb:1,ub:2047`
	Periodicity4096 *utils.INTEGER `lb:1,ub:4095`
	Periodicity8192 *utils.INTEGER `lb:1,ub:8191`
}

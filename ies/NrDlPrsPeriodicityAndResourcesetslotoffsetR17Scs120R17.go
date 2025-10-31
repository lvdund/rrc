package ies

import "rrc/utils"

// NR-DL-PRS-Periodicity-and-ResourceSetSlotOffset-r17-scs120-r17 ::= CHOICE
// Extensible
const (
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceNothing = iota
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN32R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN40R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN64R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN80R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN128R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN160R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN256R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN320R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN512R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN640R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN1280R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN2560R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN5120R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN10240R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN20480R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN40960R17
	NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17ChoiceN81920R17
)

type NrDlPrsPeriodicityAndResourcesetslotoffsetR17Scs120R17 struct {
	Choice    uint64
	N32R17    *utils.INTEGER `lb:0,ub:31`
	N40R17    *utils.INTEGER `lb:0,ub:39`
	N64R17    *utils.INTEGER `lb:0,ub:63`
	N80R17    *utils.INTEGER `lb:0,ub:79`
	N128R17   *utils.INTEGER `lb:0,ub:127`
	N160R17   *utils.INTEGER `lb:0,ub:159`
	N256R17   *utils.INTEGER `lb:0,ub:255`
	N320R17   *utils.INTEGER `lb:0,ub:319`
	N512R17   *utils.INTEGER `lb:0,ub:511`
	N640R17   *utils.INTEGER `lb:0,ub:639`
	N1280R17  *utils.INTEGER `lb:0,ub:1279`
	N2560R17  *utils.INTEGER `lb:0,ub:2559`
	N5120R17  *utils.INTEGER `lb:0,ub:5119`
	N10240R17 *utils.INTEGER `lb:0,ub:10239`
	N20480R17 *utils.INTEGER `lb:0,ub:20479`
	N40960R17 *utils.INTEGER `lb:0,ub:40959`
	N81920R17 *utils.INTEGER `lb:0,ub:81919`
}

package ies

import "rrc/utils"

// DL-PPW-PeriodicityAndStartSlot-r17-scs60 ::= CHOICE
// Extensible
const (
	DlPpwPeriodicityandstartslotR17Scs60ChoiceNothing = iota
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN16
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN20
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN32
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN40
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN64
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN80
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN128
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN160
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN256
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN320
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN640
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN1280
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN2560
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN5120
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN10240
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN20480
	DlPpwPeriodicityandstartslotR17Scs60ChoiceN40960
)

type DlPpwPeriodicityandstartslotR17Scs60 struct {
	Choice uint64
	N16    *utils.INTEGER `lb:0,ub:15`
	N20    *utils.INTEGER `lb:0,ub:19`
	N32    *utils.INTEGER `lb:0,ub:31`
	N40    *utils.INTEGER `lb:0,ub:39`
	N64    *utils.INTEGER `lb:0,ub:63`
	N80    *utils.INTEGER `lb:0,ub:79`
	N128   *utils.INTEGER `lb:0,ub:127`
	N160   *utils.INTEGER `lb:0,ub:159`
	N256   *utils.INTEGER `lb:0,ub:255`
	N320   *utils.INTEGER `lb:0,ub:319`
	N640   *utils.INTEGER `lb:0,ub:639`
	N1280  *utils.INTEGER `lb:0,ub:1279`
	N2560  *utils.INTEGER `lb:0,ub:2559`
	N5120  *utils.INTEGER `lb:0,ub:5119`
	N10240 *utils.INTEGER `lb:0,ub:10239`
	N20480 *utils.INTEGER `lb:0,ub:20479`
	N40960 *utils.INTEGER `lb:0,ub:40959`
}

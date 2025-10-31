package ies

import "rrc/utils"

// DL-PPW-PeriodicityAndStartSlot-r17-scs120 ::= CHOICE
// Extensible
const (
	DlPpwPeriodicityandstartslotR17Scs120ChoiceNothing = iota
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN32
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN40
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN64
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN80
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN128
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN160
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN256
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN320
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN512
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN640
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN1280
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN2560
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN5120
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN10240
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN20480
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN40960
	DlPpwPeriodicityandstartslotR17Scs120ChoiceN81920
)

type DlPpwPeriodicityandstartslotR17Scs120 struct {
	Choice uint64
	N32    *utils.INTEGER `lb:0,ub:31`
	N40    *utils.INTEGER `lb:0,ub:39`
	N64    *utils.INTEGER `lb:0,ub:63`
	N80    *utils.INTEGER `lb:0,ub:79`
	N128   *utils.INTEGER `lb:0,ub:127`
	N160   *utils.INTEGER `lb:0,ub:159`
	N256   *utils.INTEGER `lb:0,ub:255`
	N320   *utils.INTEGER `lb:0,ub:319`
	N512   *utils.INTEGER `lb:0,ub:511`
	N640   *utils.INTEGER `lb:0,ub:639`
	N1280  *utils.INTEGER `lb:0,ub:1279`
	N2560  *utils.INTEGER `lb:0,ub:2559`
	N5120  *utils.INTEGER `lb:0,ub:5119`
	N10240 *utils.INTEGER `lb:0,ub:10239`
	N20480 *utils.INTEGER `lb:0,ub:20479`
	N40960 *utils.INTEGER `lb:0,ub:40959`
	N81920 *utils.INTEGER `lb:0,ub:81919`
}

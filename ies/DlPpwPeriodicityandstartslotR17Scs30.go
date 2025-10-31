package ies

import "rrc/utils"

// DL-PPW-PeriodicityAndStartSlot-r17-scs30 ::= CHOICE
// Extensible
const (
	DlPpwPeriodicityandstartslotR17Scs30ChoiceNothing = iota
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN8
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN10
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN16
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN20
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN32
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN40
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN64
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN80
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN128
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN160
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN320
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN640
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN1280
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN2560
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN5120
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN10240
	DlPpwPeriodicityandstartslotR17Scs30ChoiceN20480
)

type DlPpwPeriodicityandstartslotR17Scs30 struct {
	Choice uint64
	N8     *utils.INTEGER `lb:0,ub:7`
	N10    *utils.INTEGER `lb:0,ub:9`
	N16    *utils.INTEGER `lb:0,ub:15`
	N20    *utils.INTEGER `lb:0,ub:19`
	N32    *utils.INTEGER `lb:0,ub:31`
	N40    *utils.INTEGER `lb:0,ub:39`
	N64    *utils.INTEGER `lb:0,ub:63`
	N80    *utils.INTEGER `lb:0,ub:79`
	N128   *utils.INTEGER `lb:0,ub:127`
	N160   *utils.INTEGER `lb:0,ub:159`
	N320   *utils.INTEGER `lb:0,ub:319`
	N640   *utils.INTEGER `lb:0,ub:639`
	N1280  *utils.INTEGER `lb:0,ub:1279`
	N2560  *utils.INTEGER `lb:0,ub:2559`
	N5120  *utils.INTEGER `lb:0,ub:5119`
	N10240 *utils.INTEGER `lb:0,ub:10239`
	N20480 *utils.INTEGER `lb:0,ub:20479`
}

package ies

import "rrc/utils"

// DL-PPW-PeriodicityAndStartSlot-r17-scs15 ::= CHOICE
// Extensible
const (
	DlPpwPeriodicityandstartslotR17Scs15ChoiceNothing = iota
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN4
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN5
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN8
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN10
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN16
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN20
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN32
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN40
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN64
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN80
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN160
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN320
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN640
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN1280
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN2560
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN5120
	DlPpwPeriodicityandstartslotR17Scs15ChoiceN10240
)

type DlPpwPeriodicityandstartslotR17Scs15 struct {
	Choice uint64
	N4     *utils.INTEGER `lb:0,ub:3`
	N5     *utils.INTEGER `lb:0,ub:4`
	N8     *utils.INTEGER `lb:0,ub:7`
	N10    *utils.INTEGER `lb:0,ub:9`
	N16    *utils.INTEGER `lb:0,ub:15`
	N20    *utils.INTEGER `lb:0,ub:19`
	N32    *utils.INTEGER `lb:0,ub:31`
	N40    *utils.INTEGER `lb:0,ub:39`
	N64    *utils.INTEGER `lb:0,ub:63`
	N80    *utils.INTEGER `lb:0,ub:79`
	N160   *utils.INTEGER `lb:0,ub:159`
	N320   *utils.INTEGER `lb:0,ub:319`
	N640   *utils.INTEGER `lb:0,ub:639`
	N1280  *utils.INTEGER `lb:0,ub:1279`
	N2560  *utils.INTEGER `lb:0,ub:2559`
	N5120  *utils.INTEGER `lb:0,ub:5119`
	N10240 *utils.INTEGER `lb:0,ub:10239`
}

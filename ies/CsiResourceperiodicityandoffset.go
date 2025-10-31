package ies

import "rrc/utils"

// CSI-ResourcePeriodicityAndOffset ::= CHOICE
const (
	CsiResourceperiodicityandoffsetChoiceNothing = iota
	CsiResourceperiodicityandoffsetChoiceSlots4
	CsiResourceperiodicityandoffsetChoiceSlots5
	CsiResourceperiodicityandoffsetChoiceSlots8
	CsiResourceperiodicityandoffsetChoiceSlots10
	CsiResourceperiodicityandoffsetChoiceSlots16
	CsiResourceperiodicityandoffsetChoiceSlots20
	CsiResourceperiodicityandoffsetChoiceSlots32
	CsiResourceperiodicityandoffsetChoiceSlots40
	CsiResourceperiodicityandoffsetChoiceSlots64
	CsiResourceperiodicityandoffsetChoiceSlots80
	CsiResourceperiodicityandoffsetChoiceSlots160
	CsiResourceperiodicityandoffsetChoiceSlots320
	CsiResourceperiodicityandoffsetChoiceSlots640
)

type CsiResourceperiodicityandoffset struct {
	Choice   uint64
	Slots4   *utils.INTEGER `lb:0,ub:3`
	Slots5   *utils.INTEGER `lb:0,ub:4`
	Slots8   *utils.INTEGER `lb:0,ub:7`
	Slots10  *utils.INTEGER `lb:0,ub:9`
	Slots16  *utils.INTEGER `lb:0,ub:15`
	Slots20  *utils.INTEGER `lb:0,ub:19`
	Slots32  *utils.INTEGER `lb:0,ub:31`
	Slots40  *utils.INTEGER `lb:0,ub:39`
	Slots64  *utils.INTEGER `lb:0,ub:63`
	Slots80  *utils.INTEGER `lb:0,ub:79`
	Slots160 *utils.INTEGER `lb:0,ub:159`
	Slots320 *utils.INTEGER `lb:0,ub:319`
	Slots640 *utils.INTEGER `lb:0,ub:639`
}

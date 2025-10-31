package ies

import "rrc/utils"

// CSI-ReportPeriodicityAndOffset ::= CHOICE
const (
	CsiReportperiodicityandoffsetChoiceNothing = iota
	CsiReportperiodicityandoffsetChoiceSlots4
	CsiReportperiodicityandoffsetChoiceSlots5
	CsiReportperiodicityandoffsetChoiceSlots8
	CsiReportperiodicityandoffsetChoiceSlots10
	CsiReportperiodicityandoffsetChoiceSlots16
	CsiReportperiodicityandoffsetChoiceSlots20
	CsiReportperiodicityandoffsetChoiceSlots40
	CsiReportperiodicityandoffsetChoiceSlots80
	CsiReportperiodicityandoffsetChoiceSlots160
	CsiReportperiodicityandoffsetChoiceSlots320
)

type CsiReportperiodicityandoffset struct {
	Choice   uint64
	Slots4   *utils.INTEGER `lb:0,ub:3`
	Slots5   *utils.INTEGER `lb:0,ub:4`
	Slots8   *utils.INTEGER `lb:0,ub:7`
	Slots10  *utils.INTEGER `lb:0,ub:9`
	Slots16  *utils.INTEGER `lb:0,ub:15`
	Slots20  *utils.INTEGER `lb:0,ub:19`
	Slots40  *utils.INTEGER `lb:0,ub:39`
	Slots80  *utils.INTEGER `lb:0,ub:79`
	Slots160 *utils.INTEGER `lb:0,ub:159`
	Slots320 *utils.INTEGER `lb:0,ub:319`
}

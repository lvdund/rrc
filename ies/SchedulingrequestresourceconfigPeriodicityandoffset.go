package ies

import "rrc/utils"

// SchedulingRequestResourceConfig-periodicityAndOffset ::= CHOICE
const (
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceNothing = iota
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSym2
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSym6or7
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSl1
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSl2
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSl4
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSl5
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSl8
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSl10
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSl16
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSl20
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSl40
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSl80
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSl160
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSl320
	SchedulingrequestresourceconfigPeriodicityandoffsetChoiceSl640
)

type SchedulingrequestresourceconfigPeriodicityandoffset struct {
	Choice  uint64
	Sym2    *struct{}
	Sym6or7 *struct{}
	Sl1     *struct{}
	Sl2     *utils.INTEGER `lb:0,ub:1`
	Sl4     *utils.INTEGER `lb:0,ub:3`
	Sl5     *utils.INTEGER `lb:0,ub:4`
	Sl8     *utils.INTEGER `lb:0,ub:7`
	Sl10    *utils.INTEGER `lb:0,ub:9`
	Sl16    *utils.INTEGER `lb:0,ub:15`
	Sl20    *utils.INTEGER `lb:0,ub:19`
	Sl40    *utils.INTEGER `lb:0,ub:39`
	Sl80    *utils.INTEGER `lb:0,ub:79`
	Sl160   *utils.INTEGER `lb:0,ub:159`
	Sl320   *utils.INTEGER `lb:0,ub:319`
	Sl640   *utils.INTEGER `lb:0,ub:639`
}

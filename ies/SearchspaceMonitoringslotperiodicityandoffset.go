package ies

import "rrc/utils"

// SearchSpace-monitoringSlotPeriodicityAndOffset ::= CHOICE
const (
	SearchspaceMonitoringslotperiodicityandoffsetChoiceNothing = iota
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl1
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl2
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl4
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl5
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl8
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl10
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl16
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl20
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl40
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl80
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl160
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl320
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl640
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl1280
	SearchspaceMonitoringslotperiodicityandoffsetChoiceSl2560
)

type SearchspaceMonitoringslotperiodicityandoffset struct {
	Choice uint64
	Sl1    *struct{}
	Sl2    *utils.INTEGER `lb:0,ub:1`
	Sl4    *utils.INTEGER `lb:0,ub:3`
	Sl5    *utils.INTEGER `lb:0,ub:4`
	Sl8    *utils.INTEGER `lb:0,ub:7`
	Sl10   *utils.INTEGER `lb:0,ub:9`
	Sl16   *utils.INTEGER `lb:0,ub:15`
	Sl20   *utils.INTEGER `lb:0,ub:19`
	Sl40   *utils.INTEGER `lb:0,ub:39`
	Sl80   *utils.INTEGER `lb:0,ub:79`
	Sl160  *utils.INTEGER `lb:0,ub:159`
	Sl320  *utils.INTEGER `lb:0,ub:319`
	Sl640  *utils.INTEGER `lb:0,ub:639`
	Sl1280 *utils.INTEGER `lb:0,ub:1279`
	Sl2560 *utils.INTEGER `lb:0,ub:2559`
}

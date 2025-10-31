package ies

import "rrc/utils"

// SearchSpaceExt-v1700-monitoringSlotPeriodicityAndOffset-v1710 ::= CHOICE
const (
	SearchspaceextV1700MonitoringslotperiodicityandoffsetV1710ChoiceNothing = iota
	SearchspaceextV1700MonitoringslotperiodicityandoffsetV1710ChoiceSl32
	SearchspaceextV1700MonitoringslotperiodicityandoffsetV1710ChoiceSl64
	SearchspaceextV1700MonitoringslotperiodicityandoffsetV1710ChoiceSl128
	SearchspaceextV1700MonitoringslotperiodicityandoffsetV1710ChoiceSl5120
	SearchspaceextV1700MonitoringslotperiodicityandoffsetV1710ChoiceSl10240
	SearchspaceextV1700MonitoringslotperiodicityandoffsetV1710ChoiceSl20480
)

type SearchspaceextV1700MonitoringslotperiodicityandoffsetV1710 struct {
	Choice  uint64
	Sl32    *utils.INTEGER `lb:0,ub:31`
	Sl64    *utils.INTEGER `lb:0,ub:63`
	Sl128   *utils.INTEGER `lb:0,ub:127`
	Sl5120  *utils.INTEGER `lb:0,ub:5119`
	Sl10240 *utils.INTEGER `lb:0,ub:10239`
	Sl20480 *utils.INTEGER `lb:0,ub:20479`
}

package ies

import "rrc/utils"

// SearchSpaceExt-v1700-monitoringSlotsWithinSlotGroup-r17 ::= CHOICE
const (
	SearchspaceextV1700MonitoringslotswithinslotgroupR17ChoiceNothing = iota
	SearchspaceextV1700MonitoringslotswithinslotgroupR17ChoiceSlotgrouplength4R17
	SearchspaceextV1700MonitoringslotswithinslotgroupR17ChoiceSlotgrouplength8R17
)

type SearchspaceextV1700MonitoringslotswithinslotgroupR17 struct {
	Choice              uint64
	Slotgrouplength4R17 *utils.BITSTRING `lb:4,ub:4`
	Slotgrouplength8R17 *utils.BITSTRING `lb:8,ub:8`
}

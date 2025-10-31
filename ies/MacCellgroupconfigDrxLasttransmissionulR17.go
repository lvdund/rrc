package ies

import "rrc/utils"

// MAC-CellGroupConfig-drx-LastTransmissionUL-r17 ::= ENUMERATED
type MacCellgroupconfigDrxLasttransmissionulR17 struct {
	Value utils.ENUMERATED
}

const (
	MacCellgroupconfigDrxLasttransmissionulR17EnumeratedNothing = iota
	MacCellgroupconfigDrxLasttransmissionulR17EnumeratedEnabled
)

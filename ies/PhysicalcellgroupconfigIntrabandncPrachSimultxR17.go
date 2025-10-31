package ies

import "rrc/utils"

// PhysicalCellGroupConfig-intraBandNC-PRACH-simulTx-r17 ::= ENUMERATED
type PhysicalcellgroupconfigIntrabandncPrachSimultxR17 struct {
	Value utils.ENUMERATED
}

const (
	PhysicalcellgroupconfigIntrabandncPrachSimultxR17EnumeratedNothing = iota
	PhysicalcellgroupconfigIntrabandncPrachSimultxR17EnumeratedEnabled
)

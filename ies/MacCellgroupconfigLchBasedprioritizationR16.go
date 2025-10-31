package ies

import "rrc/utils"

// MAC-CellGroupConfig-lch-BasedPrioritization-r16 ::= ENUMERATED
type MacCellgroupconfigLchBasedprioritizationR16 struct {
	Value utils.ENUMERATED
}

const (
	MacCellgroupconfigLchBasedprioritizationR16EnumeratedNothing = iota
	MacCellgroupconfigLchBasedprioritizationR16EnumeratedEnabled
)

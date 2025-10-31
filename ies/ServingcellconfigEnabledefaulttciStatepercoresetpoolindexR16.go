package ies

import "rrc/utils"

// ServingCellConfig-enableDefaultTCI-StatePerCoresetPoolIndex-r16 ::= ENUMERATED
type ServingcellconfigEnabledefaulttciStatepercoresetpoolindexR16 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigEnabledefaulttciStatepercoresetpoolindexR16EnumeratedNothing = iota
	ServingcellconfigEnabledefaulttciStatepercoresetpoolindexR16EnumeratedEnabled
)

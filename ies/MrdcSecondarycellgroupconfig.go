package ies

import "rrc/utils"

// MRDC-SecondaryCellGroupConfig ::= SEQUENCE
type MrdcSecondarycellgroupconfig struct {
	MrdcReleaseandadd      *utils.BOOLEAN
	MrdcSecondarycellgroup MrdcSecondarycellgroupconfigMrdcSecondarycellgroup
}

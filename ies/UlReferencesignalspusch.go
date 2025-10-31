package ies

import "rrc/utils"

// UL-ReferenceSignalsPUSCH ::= SEQUENCE
type UlReferencesignalspusch struct {
	Grouphoppingenabled    utils.BOOLEAN
	Groupassignmentpusch   utils.INTEGER `lb:0,ub:29`
	Sequencehoppingenabled utils.BOOLEAN
	Cyclicshift            utils.INTEGER `lb:0,ub:7`
}

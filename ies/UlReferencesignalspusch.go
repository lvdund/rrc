package ies

import "rrc/utils"

// UL-ReferenceSignalsPUSCH ::= SEQUENCE
type UlReferencesignalspusch struct {
	Grouphoppingenabled    bool
	Groupassignmentpusch   utils.INTEGER
	Sequencehoppingenabled bool
	Cyclicshift            utils.INTEGER
}

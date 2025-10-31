package ies

import "rrc/utils"

// SubgroupConfig-r17 ::= SEQUENCE
// Extensible
type SubgroupconfigR17 struct {
	SubgroupsnumperpoR17   utils.INTEGER  `lb:0,ub:maxNrofPagingSubgroupsR17`
	SubgroupsnumforueidR17 *utils.INTEGER `lb:0,ub:maxNrofPagingSubgroupsR17`
}

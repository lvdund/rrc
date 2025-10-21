package ies

import "rrc/utils"

// RedistributionNeighCellList-r13 ::= SEQUENCE OF RedistributionNeighCell-r13
// SIZE (1..maxCellInter)
type RedistributionneighcelllistR13 struct {
	Value []RedistributionneighcellR13
}

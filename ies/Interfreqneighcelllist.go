package ies

import "rrc/utils"

// InterFreqNeighCellList ::= SEQUENCE OF InterFreqNeighCellInfo
// SIZE (1..maxCellInter)
type Interfreqneighcelllist struct {
	Value []Interfreqneighcellinfo
}

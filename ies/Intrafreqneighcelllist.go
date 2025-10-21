package ies

import "rrc/utils"

// IntraFreqNeighCellList ::= SEQUENCE OF IntraFreqNeighCellInfo
// SIZE (1..maxCellIntra)
type Intrafreqneighcelllist struct {
	Value utils.Sequence[Intrafreqneighcellinfo]
}

package ies

import "rrc/utils"

// IntraFreqNeighCellList-NB-v1530 ::= SEQUENCE OF IntraFreqNeighCellInfo-NB-v1530
// SIZE (1..maxCellIntra)
type IntrafreqneighcelllistNbV1530 struct {
	Value []IntrafreqneighcellinfoNbV1530
}

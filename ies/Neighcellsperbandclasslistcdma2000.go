package ies

import "rrc/utils"

// NeighCellsPerBandclassListCDMA2000 ::= SEQUENCE OF NeighCellsPerBandclassCDMA2000
// SIZE (1..16)
type Neighcellsperbandclasslistcdma2000 struct {
	Value utils.Sequence[Neighcellsperbandclasscdma2000]
}

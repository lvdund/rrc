package ies

import "rrc/utils"

// NeighCellsPerBandclassListCDMA2000-v920 ::= SEQUENCE OF NeighCellsPerBandclassCDMA2000-v920
// SIZE (1..16)
type Neighcellsperbandclasslistcdma2000V920 struct {
	Value utils.Sequence[Neighcellsperbandclasscdma2000V920]
}

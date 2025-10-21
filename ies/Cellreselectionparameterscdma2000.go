package ies

import "rrc/utils"

// CellReselectionParametersCDMA2000 ::= SEQUENCE
type Cellreselectionparameterscdma2000 struct {
	Bandclasslist          Bandclasslistcdma2000
	Neighcelllist          Neighcelllistcdma2000
	TReselectioncdma2000   TReselection
	TReselectioncdma2000Sf *Speedstatescalefactors
}

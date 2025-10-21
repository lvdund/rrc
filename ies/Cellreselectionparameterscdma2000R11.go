package ies

import "rrc/utils"

// CellReselectionParametersCDMA2000-r11 ::= SEQUENCE
type Cellreselectionparameterscdma2000R11 struct {
	Bandclasslist          Bandclasslistcdma2000
	NeighcelllistR11       interface{}
	TReselectioncdma2000   TReselection
	TReselectioncdma2000Sf *Speedstatescalefactors
}

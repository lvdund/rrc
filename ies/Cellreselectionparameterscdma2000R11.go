package ies

// CellReselectionParametersCDMA2000-r11 ::= SEQUENCE
type Cellreselectionparameterscdma2000R11 struct {
	Bandclasslist          Bandclasslistcdma2000
	NeighcelllistR11       []Neighcellcdma2000R11 `lb:1,ub:16`
	TReselectioncdma2000   TReselection
	TReselectioncdma2000Sf *Speedstatescalefactors
}

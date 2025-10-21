package ies

import "rrc/utils"

// NeighCellsCRS-Info-r15 ::= CHOICE
type NeighcellscrsInfoR15 interface {
	isNeighcellscrsInfoR15()
}

type NeighcellscrsInfoR15Release struct {
	Value struct{}
}

func (*NeighcellscrsInfoR15Release) isNeighcellscrsInfoR15() {}

type NeighcellscrsInfoR15Setup struct {
	Value CrsAssistanceinfolistR15
}

func (*NeighcellscrsInfoR15Setup) isNeighcellscrsInfoR15() {}

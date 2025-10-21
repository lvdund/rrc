package ies

import "rrc/utils"

// NeighCellsCRS-Info-r13 ::= CHOICE
type NeighcellscrsInfoR13 interface {
	isNeighcellscrsInfoR13()
}

type NeighcellscrsInfoR13Release struct {
	Value struct{}
}

func (*NeighcellscrsInfoR13Release) isNeighcellscrsInfoR13() {}

type NeighcellscrsInfoR13Setup struct {
	Value CrsAssistanceinfolistR13
}

func (*NeighcellscrsInfoR13Setup) isNeighcellscrsInfoR13() {}

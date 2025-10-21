package ies

import "rrc/utils"

// NeighCellsCRS-Info-r11 ::= CHOICE
type NeighcellscrsInfoR11 interface {
	isNeighcellscrsInfoR11()
}

type NeighcellscrsInfoR11Release struct {
	Value struct{}
}

func (*NeighcellscrsInfoR11Release) isNeighcellscrsInfoR11() {}

type NeighcellscrsInfoR11Setup struct {
	Value CrsAssistanceinfolistR11
}

func (*NeighcellscrsInfoR11Setup) isNeighcellscrsInfoR11() {}

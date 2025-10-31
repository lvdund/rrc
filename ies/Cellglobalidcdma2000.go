package ies

import "rrc/utils"

// CellGlobalIdCDMA2000 ::= CHOICE
const (
	Cellglobalidcdma2000ChoiceNothing = iota
	Cellglobalidcdma2000ChoiceCellglobalid1xrtt
	Cellglobalidcdma2000ChoiceCellglobalidhrpd
)

type Cellglobalidcdma2000 struct {
	Choice            uint64
	Cellglobalid1xrtt *utils.BITSTRING `lb:47,ub:47`
	Cellglobalidhrpd  *utils.BITSTRING `lb:128,ub:128`
}

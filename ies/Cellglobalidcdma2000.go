package ies

import "rrc/utils"

// CellGlobalIdCDMA2000 ::= CHOICE
type Cellglobalidcdma2000 interface {
	isCellglobalidcdma2000()
}

type Cellglobalidcdma2000Cellglobalid1xrtt struct {
	Value utils.BITSTRING
}

func (*Cellglobalidcdma2000Cellglobalid1xrtt) isCellglobalidcdma2000() {}

type Cellglobalidcdma2000Cellglobalidhrpd struct {
	Value utils.BITSTRING
}

func (*Cellglobalidcdma2000Cellglobalidhrpd) isCellglobalidcdma2000() {}

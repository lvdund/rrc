package ies

import "rrc/utils"

// BandNR-condPSCellChange-r16 ::= ENUMERATED
type BandnrCondpscellchangeR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrCondpscellchangeR16EnumeratedNothing = iota
	BandnrCondpscellchangeR16EnumeratedSupported
)

package ies

import "rrc/utils"

// BandSidelinkPC5-r16-scheme2-ConflictDeterminationRSRP-r17 ::= ENUMERATED
type Bandsidelinkpc5R16Scheme2ConflictdeterminationrsrpR17 struct {
	Value utils.ENUMERATED
}

const (
	Bandsidelinkpc5R16Scheme2ConflictdeterminationrsrpR17EnumeratedNothing = iota
	Bandsidelinkpc5R16Scheme2ConflictdeterminationrsrpR17EnumeratedSupported
)

package ies

// P0-PUSCH-Set-r16 ::= SEQUENCE
// Extensible
type P0PuschSetR16 struct {
	P0PuschSetidR16 P0PuschSetidR16
	P0ListR16       *[]P0PuschR16 `lb:1,ub:maxNrofP0PuschSetR16`
}

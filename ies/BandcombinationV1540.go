package ies

// BandCombination-v1540 ::= SEQUENCE
type BandcombinationV1540 struct {
	BandlistV1540       []BandparametersV1540 `lb:1,ub:maxSimultaneousBands`
	CaParametersnrV1540 *CaParametersnrV1540
}

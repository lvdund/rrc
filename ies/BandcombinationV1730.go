package ies

// BandCombination-v1730 ::= SEQUENCE
type BandcombinationV1730 struct {
	CaParametersnrV1730   *CaParametersnrV1730
	CaParametersnrdcV1730 *CaParametersnrdcV1730
	BandlistV1730         *[]BandparametersV1730 `lb:1,ub:maxSimultaneousBands`
}

package ies

// BandCombination-v1560 ::= SEQUENCE
type BandcombinationV1560 struct {
	NeDcBc                 *BandcombinationV1560NeDcBc
	CaParametersnrdc       *CaParametersnrdc
	CaParameterseutraV1560 *CaParameterseutraV1560
	CaParametersnrV1560    *CaParametersnrV1560
}

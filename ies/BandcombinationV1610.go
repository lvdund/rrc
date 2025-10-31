package ies

// BandCombination-v1610 ::= SEQUENCE
type BandcombinationV1610 struct {
	BandlistV1610                *[]BandparametersV1610 `lb:1,ub:maxSimultaneousBands`
	CaParametersnrV1610          *CaParametersnrV1610
	CaParametersnrdcV1610        *CaParametersnrdcV1610
	PowerclassV1610              *BandcombinationV1610PowerclassV1610
	PowerclassnrpartR16          *BandcombinationV1610PowerclassnrpartR16
	FeaturesetcombinationdapsR16 *Featuresetcombinationid
	MrdcParametersV1620          *MrdcParametersV1620
}

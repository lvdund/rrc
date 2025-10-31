package ies

// BandCombinationParameters-v1320 ::= SEQUENCE
type BandcombinationparametersV1320 struct {
	BandparameterlistV1320          *[]BandparametersV1320 `lb:1,ub:maxSimultaneousBandsR10`
	AdditionalrxTxPerformancereqR13 *BandcombinationparametersV1320AdditionalrxTxPerformancereqR13
}

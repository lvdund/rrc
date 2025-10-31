package ies

// BandCombinationParameters-v1530 ::= SEQUENCE
type BandcombinationparametersV1530 struct {
	BandparameterlistV1530 *[]BandparametersV1530 `lb:1,ub:maxSimultaneousBandsR10`
	SptParametersR15       *SptParametersR15
}

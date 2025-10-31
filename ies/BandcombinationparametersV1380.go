package ies

// BandCombinationParameters-v1380 ::= SEQUENCE
type BandcombinationparametersV1380 struct {
	BandparameterlistV1380 *[]BandparametersV1380 `lb:1,ub:maxSimultaneousBandsR10`
}

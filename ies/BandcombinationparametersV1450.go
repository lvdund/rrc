package ies

// BandCombinationParameters-v1450 ::= SEQUENCE
type BandcombinationparametersV1450 struct {
	BandparameterlistV1450 *[]BandparametersV1450 `lb:1,ub:maxSimultaneousBandsR10`
}

package ies

// BandCombinationParameters-v1270 ::= SEQUENCE
type BandcombinationparametersV1270 struct {
	BandparameterlistV1270 *[]BandparametersV1270 `lb:1,ub:maxSimultaneousBandsR10`
}

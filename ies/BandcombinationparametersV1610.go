package ies

// BandCombinationParameters-v1610 ::= SEQUENCE
type BandcombinationparametersV1610 struct {
	MeasgapinfonrR16       *MeasgapinfonrR16
	BandparameterlistV1610 *[]BandparametersV1610 `lb:1,ub:maxSimultaneousBandsR10`
	InterfreqdapsR16       *BandcombinationparametersV1610InterfreqdapsR16
}

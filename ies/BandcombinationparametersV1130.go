package ies

// BandCombinationParameters-v1130 ::= SEQUENCE
// Extensible
type BandcombinationparametersV1130 struct {
	MultipletimingadvanceR11 *BandcombinationparametersV1130MultipletimingadvanceR11
	SimultaneousrxTxR11      *BandcombinationparametersV1130SimultaneousrxTxR11
	BandparameterlistR11     *[]BandparametersV1130 `lb:1,ub:maxSimultaneousBandsR10`
}

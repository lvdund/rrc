package ies

// BandCombinationParameters-r11 ::= SEQUENCE
// Extensible
type BandcombinationparametersR11 struct {
	BandparameterlistR11                []BandparametersR11 `lb:1,ub:maxSimultaneousBandsR10`
	SupportedbandwidthcombinationsetR11 *SupportedbandwidthcombinationsetR10
	MultipletimingadvanceR11            *BandcombinationparametersR11MultipletimingadvanceR11
	SimultaneousrxTxR11                 *BandcombinationparametersR11SimultaneousrxTxR11
	BandinfoeutraR11                    Bandinfoeutra
}

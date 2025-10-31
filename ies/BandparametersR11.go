package ies

// BandParameters-r11 ::= SEQUENCE
type BandparametersR11 struct {
	BandeutraR11        FreqbandindicatorR11
	BandparametersulR11 *BandparametersulR10
	BandparametersdlR11 *BandparametersdlR10
	SupportedcsiProcR11 *BandparametersR11SupportedcsiProcR11
}

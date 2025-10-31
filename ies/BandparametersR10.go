package ies

// BandParameters-r10 ::= SEQUENCE
type BandparametersR10 struct {
	BandeutraR10        Freqbandindicator
	BandparametersulR10 *BandparametersulR10
	BandparametersdlR10 *BandparametersdlR10
}

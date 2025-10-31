package ies

// V2X-BandParameters-r14 ::= SEQUENCE
type V2xBandparametersR14 struct {
	V2xFreqbandeutraR14   FreqbandindicatorR11
	BandparameterstxslR14 *BandparameterstxslR14
	BandparametersrxslR14 *BandparametersrxslR14
}

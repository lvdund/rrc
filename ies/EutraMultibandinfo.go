package ies

// EUTRA-MultiBandInfo ::= SEQUENCE
type EutraMultibandinfo struct {
	EutraFreqbandindicator Freqbandindicatoreutra
	EutraNsPmaxlist        *EutraNsPmaxlist
}

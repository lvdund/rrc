package ies

// MeasResult3EUTRA-r15 ::= SEQUENCE
// Extensible
type Measresult3eutraR15 struct {
	CarrierfreqR15             ArfcnValueeutraR9
	MeasresultservingcellR15   *Measresulteutra
	MeasresultneighcelllistR15 *Measresultlisteutra
}

package ies

// BandParameters-eutra ::= SEQUENCE
type BandparametersEutra struct {
	Bandeutra               Freqbandindicatoreutra
	CaBandwidthclassdlEutra *CaBandwidthclasseutra
	CaBandwidthclassulEutra *CaBandwidthclasseutra
}

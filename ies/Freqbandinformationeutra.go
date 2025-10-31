package ies

// FreqBandInformationEUTRA ::= SEQUENCE
type Freqbandinformationeutra struct {
	Bandeutra               Freqbandindicatoreutra
	CaBandwidthclassdlEutra *CaBandwidthclasseutra
	CaBandwidthclassulEutra *CaBandwidthclasseutra
}

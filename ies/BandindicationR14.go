package ies

// BandIndication-r14 ::= SEQUENCE
type BandindicationR14 struct {
	BandeutraR14          FreqbandindicatorR11
	CaBandwidthclassdlR14 CaBandwidthclassR10
	CaBandwidthclassulR14 *CaBandwidthclassR10
}

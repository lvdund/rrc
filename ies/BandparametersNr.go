package ies

// BandParameters-nr ::= SEQUENCE
type BandparametersNr struct {
	Bandnr               Freqbandindicatornr
	CaBandwidthclassdlNr *CaBandwidthclassnr
	CaBandwidthclassulNr *CaBandwidthclassnr
}

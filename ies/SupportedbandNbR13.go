package ies

// SupportedBand-NB-r13 ::= SEQUENCE
type SupportedbandNbR13 struct {
	BandR13              FreqbandindicatorNbR13
	Powerclassnb20dbmR13 *SupportedbandNbR13Powerclassnb20dbmR13
}

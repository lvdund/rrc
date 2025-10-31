package ies

// BandInfoEUTRA ::= SEQUENCE
type Bandinfoeutra struct {
	Interfreqbandlist Interfreqbandlist
	InterratBandlist  *InterratBandlist
}

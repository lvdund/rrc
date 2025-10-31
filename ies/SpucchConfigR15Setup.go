package ies

// SPUCCH-Config-r15-setup ::= SEQUENCE
type SpucchConfigR15Setup struct {
	SpucchSetR15                               *SpucchSetR15
	TwoantennaportactivatedspucchFormat1a1bR15 *bool
	Dummy                                      SpucchConfigR15SetupDummy
}

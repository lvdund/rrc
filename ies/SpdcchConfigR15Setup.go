package ies

// SPDCCH-Config-r15-setup ::= SEQUENCE
type SpdcchConfigR15Setup struct {
	SpdcchL1ReuseindicationR15 *SpdcchConfigR15SetupSpdcchL1ReuseindicationR15
	SpdcchSetconfigR15         *SpdcchSetR15
}

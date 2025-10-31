package ies

import "rrc/utils"

// SPDCCH-Config-r15-setup-spdcch-L1-ReuseIndication-r15 ::= ENUMERATED
type SpdcchConfigR15SetupSpdcchL1ReuseindicationR15 struct {
	Value utils.ENUMERATED
}

const (
	SpdcchConfigR15SetupSpdcchL1ReuseindicationR15EnumeratedNothing = iota
	SpdcchConfigR15SetupSpdcchL1ReuseindicationR15EnumeratedN0
	SpdcchConfigR15SetupSpdcchL1ReuseindicationR15EnumeratedN1
	SpdcchConfigR15SetupSpdcchL1ReuseindicationR15EnumeratedN2
)

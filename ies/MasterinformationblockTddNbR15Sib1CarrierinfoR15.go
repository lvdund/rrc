package ies

import "rrc/utils"

// MasterInformationBlock-TDD-NB-r15-sib1-CarrierInfo-r15 ::= ENUMERATED
type MasterinformationblockTddNbR15Sib1CarrierinfoR15 struct {
	Value utils.ENUMERATED
}

const (
	MasterinformationblockTddNbR15Sib1CarrierinfoR15EnumeratedNothing = iota
	MasterinformationblockTddNbR15Sib1CarrierinfoR15EnumeratedAnchor
	MasterinformationblockTddNbR15Sib1CarrierinfoR15EnumeratedNon_Anchor
)

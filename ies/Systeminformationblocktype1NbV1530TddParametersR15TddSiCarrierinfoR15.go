package ies

import "rrc/utils"

// SystemInformationBlockType1-NB-v1530-tdd-Parameters-r15-tdd-SI-CarrierInfo-r15 ::= ENUMERATED
type Systeminformationblocktype1NbV1530TddParametersR15TddSiCarrierinfoR15 struct {
	Value utils.ENUMERATED
}

const (
	Systeminformationblocktype1NbV1530TddParametersR15TddSiCarrierinfoR15EnumeratedNothing = iota
	Systeminformationblocktype1NbV1530TddParametersR15TddSiCarrierinfoR15EnumeratedAnchor
	Systeminformationblocktype1NbV1530TddParametersR15TddSiCarrierinfoR15EnumeratedNon_Anchor
)

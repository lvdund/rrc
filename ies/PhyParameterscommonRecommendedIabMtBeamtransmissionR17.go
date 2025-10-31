package ies

import "rrc/utils"

// Phy-ParametersCommon-recommended-IAB-MT-BeamTransmission-r17 ::= ENUMERATED
type PhyParameterscommonRecommendedIabMtBeamtransmissionR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonRecommendedIabMtBeamtransmissionR17EnumeratedNothing = iota
	PhyParameterscommonRecommendedIabMtBeamtransmissionR17EnumeratedSupported
)

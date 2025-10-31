package ies

import "rrc/utils"

// Phy-ParametersCommon-sps-HARQ-ACK-Deferral-r17-sharedSpectrumChAccess-r17 ::= ENUMERATED
type PhyParameterscommonSpsHarqAckDeferralR17SharedspectrumchaccessR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonSpsHarqAckDeferralR17SharedspectrumchaccessR17EnumeratedNothing = iota
	PhyParameterscommonSpsHarqAckDeferralR17SharedspectrumchaccessR17EnumeratedSupported
)

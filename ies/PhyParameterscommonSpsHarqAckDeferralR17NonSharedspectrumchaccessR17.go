package ies

import "rrc/utils"

// Phy-ParametersCommon-sps-HARQ-ACK-Deferral-r17-non-SharedSpectrumChAccess-r17 ::= ENUMERATED
type PhyParameterscommonSpsHarqAckDeferralR17NonSharedspectrumchaccessR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonSpsHarqAckDeferralR17NonSharedspectrumchaccessR17EnumeratedNothing = iota
	PhyParameterscommonSpsHarqAckDeferralR17NonSharedspectrumchaccessR17EnumeratedSupported
)

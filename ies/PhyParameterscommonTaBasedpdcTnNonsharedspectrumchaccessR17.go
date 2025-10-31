package ies

import "rrc/utils"

// Phy-ParametersCommon-ta-BasedPDC-TN-NonSharedSpectrumChAccess-r17 ::= ENUMERATED
type PhyParameterscommonTaBasedpdcTnNonsharedspectrumchaccessR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonTaBasedpdcTnNonsharedspectrumchaccessR17EnumeratedNothing = iota
	PhyParameterscommonTaBasedpdcTnNonsharedspectrumchaccessR17EnumeratedSupported
)

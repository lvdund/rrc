package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-ss-SINR-Meas-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16SsSinrMeasR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16SsSinrMeasR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16SsSinrMeasR16EnumeratedSupported
)

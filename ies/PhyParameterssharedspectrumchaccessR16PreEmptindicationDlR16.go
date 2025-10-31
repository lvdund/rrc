package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-pre-EmptIndication-DL-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16PreEmptindicationDlR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16PreEmptindicationDlR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16PreEmptindicationDlR16EnumeratedSupported
)

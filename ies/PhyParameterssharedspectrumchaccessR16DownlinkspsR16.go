package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-downlinkSPS-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16DownlinkspsR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16DownlinkspsR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16DownlinkspsR16EnumeratedSupported
)

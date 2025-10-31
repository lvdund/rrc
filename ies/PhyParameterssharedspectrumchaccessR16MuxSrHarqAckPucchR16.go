package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-mux-SR-HARQ-ACK-PUCCH-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16MuxSrHarqAckPucchR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16MuxSrHarqAckPucchR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16MuxSrHarqAckPucchR16EnumeratedSupported
)

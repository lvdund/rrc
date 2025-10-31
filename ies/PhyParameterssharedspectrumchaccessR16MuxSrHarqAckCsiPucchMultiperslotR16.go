package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-mux-SR-HARQ-ACK-CSI-PUCCH-MultiPerSlot-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16MuxSrHarqAckCsiPucchMultiperslotR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16MuxSrHarqAckCsiPucchMultiperslotR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16MuxSrHarqAckCsiPucchMultiperslotR16EnumeratedSupported
)

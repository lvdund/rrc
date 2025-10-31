package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-mux-SR-HARQ-ACK-CSI-PUCCH-OncePerSlot-r16-sameSymbol-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16MuxSrHarqAckCsiPucchOnceperslotR16SamesymbolR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16MuxSrHarqAckCsiPucchOnceperslotR16SamesymbolR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16MuxSrHarqAckCsiPucchOnceperslotR16SamesymbolR16EnumeratedSupported
)

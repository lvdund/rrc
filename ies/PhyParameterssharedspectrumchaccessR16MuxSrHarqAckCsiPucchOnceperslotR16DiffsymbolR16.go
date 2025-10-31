package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-mux-SR-HARQ-ACK-CSI-PUCCH-OncePerSlot-r16-diffSymbol-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16MuxSrHarqAckCsiPucchOnceperslotR16DiffsymbolR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16MuxSrHarqAckCsiPucchOnceperslotR16DiffsymbolR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16MuxSrHarqAckCsiPucchOnceperslotR16DiffsymbolR16EnumeratedSupported
)

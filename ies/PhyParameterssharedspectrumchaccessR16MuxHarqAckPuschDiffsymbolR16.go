package ies

import "rrc/utils"

// Phy-ParametersSharedSpectrumChAccess-r16-mux-HARQ-ACK-PUSCH-DiffSymbol-r16 ::= ENUMERATED
type PhyParameterssharedspectrumchaccessR16MuxHarqAckPuschDiffsymbolR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterssharedspectrumchaccessR16MuxHarqAckPuschDiffsymbolR16EnumeratedNothing = iota
	PhyParameterssharedspectrumchaccessR16MuxHarqAckPuschDiffsymbolR16EnumeratedSupported
)

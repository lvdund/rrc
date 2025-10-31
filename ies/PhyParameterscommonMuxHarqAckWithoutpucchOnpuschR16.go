package ies

import "rrc/utils"

// Phy-ParametersCommon-mux-HARQ-ACK-withoutPUCCH-onPUSCH-r16 ::= ENUMERATED
type PhyParameterscommonMuxHarqAckWithoutpucchOnpuschR16 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonMuxHarqAckWithoutpucchOnpuschR16EnumeratedNothing = iota
	PhyParameterscommonMuxHarqAckWithoutpucchOnpuschR16EnumeratedSupported
)

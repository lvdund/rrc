package ies

import "rrc/utils"

// FeatureSetUplink-v1610-mux-SR-HARQ-ACK-r16 ::= ENUMERATED
type FeaturesetuplinkV1610MuxSrHarqAckR16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610MuxSrHarqAckR16EnumeratedNothing = iota
	FeaturesetuplinkV1610MuxSrHarqAckR16EnumeratedSupported
)

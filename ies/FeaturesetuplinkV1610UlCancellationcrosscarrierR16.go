package ies

import "rrc/utils"

// FeatureSetUplink-v1610-ul-CancellationCrossCarrier-r16 ::= ENUMERATED
type FeaturesetuplinkV1610UlCancellationcrosscarrierR16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610UlCancellationcrosscarrierR16EnumeratedNothing = iota
	FeaturesetuplinkV1610UlCancellationcrosscarrierR16EnumeratedSupported
)

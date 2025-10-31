package ies

import "rrc/utils"

// FeatureSetUplink-v1610-ul-CancellationSelfCarrier-r16 ::= ENUMERATED
type FeaturesetuplinkV1610UlCancellationselfcarrierR16 struct {
	Value utils.ENUMERATED
}

const (
	FeaturesetuplinkV1610UlCancellationselfcarrierR16EnumeratedNothing = iota
	FeaturesetuplinkV1610UlCancellationselfcarrierR16EnumeratedSupported
)

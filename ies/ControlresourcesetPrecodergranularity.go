package ies

import "rrc/utils"

// ControlResourceSet-precoderGranularity ::= ENUMERATED
type ControlresourcesetPrecodergranularity struct {
	Value utils.ENUMERATED
}

const (
	ControlresourcesetPrecodergranularityEnumeratedNothing = iota
	ControlresourcesetPrecodergranularityEnumeratedSameasreg_Bundle
	ControlresourcesetPrecodergranularityEnumeratedAllcontiguousrbs
)

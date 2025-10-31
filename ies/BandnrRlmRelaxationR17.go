package ies

import "rrc/utils"

// BandNR-rlm-Relaxation-r17 ::= ENUMERATED
type BandnrRlmRelaxationR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrRlmRelaxationR17EnumeratedNothing = iota
	BandnrRlmRelaxationR17EnumeratedSupported
)

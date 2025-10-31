package ies

import "rrc/utils"

// PTRS-UplinkConfig-transformPrecoderEnabled-timeDensityTransformPrecoding ::= ENUMERATED
type PtrsUplinkconfigTransformprecoderenabledTimedensitytransformprecoding struct {
	Value utils.ENUMERATED
}

const (
	PtrsUplinkconfigTransformprecoderenabledTimedensitytransformprecodingEnumeratedNothing = iota
	PtrsUplinkconfigTransformprecoderenabledTimedensitytransformprecodingEnumeratedD2
)

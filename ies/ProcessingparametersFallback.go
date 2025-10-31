package ies

import "rrc/utils"

// ProcessingParameters-fallback ::= ENUMERATED
type ProcessingparametersFallback struct {
	Value utils.ENUMERATED
}

const (
	ProcessingparametersFallbackEnumeratedNothing = iota
	ProcessingparametersFallbackEnumeratedSc
	ProcessingparametersFallbackEnumeratedCap1_Only
)

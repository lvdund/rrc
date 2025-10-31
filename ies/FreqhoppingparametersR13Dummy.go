package ies

import "rrc/utils"

// FreqHoppingParameters-r13-dummy ::= ENUMERATED
type FreqhoppingparametersR13Dummy struct {
	Value utils.ENUMERATED
}

const (
	FreqhoppingparametersR13DummyEnumeratedNothing = iota
	FreqhoppingparametersR13DummyEnumeratedNb2
	FreqhoppingparametersR13DummyEnumeratedNb4
)

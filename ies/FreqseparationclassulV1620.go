package ies

import "rrc/utils"

// FreqSeparationClassUL-v1620 ::= ENUMERATED
type FreqseparationclassulV1620 struct {
	Value utils.ENUMERATED
}

const (
	FreqseparationclassulV1620EnumeratedNothing = iota
	FreqseparationclassulV1620EnumeratedMhz1000
)

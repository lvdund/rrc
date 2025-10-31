package ies

import "rrc/utils"

// SRS-Resource-repetitionFactor-v1730 ::= ENUMERATED
type SrsResourceRepetitionfactorV1730 struct {
	Value utils.ENUMERATED
}

const (
	SrsResourceRepetitionfactorV1730EnumeratedNothing = iota
	SrsResourceRepetitionfactorV1730EnumeratedN3
)

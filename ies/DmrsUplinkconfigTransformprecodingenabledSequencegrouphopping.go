package ies

import "rrc/utils"

// DMRS-UplinkConfig-transformPrecodingEnabled-sequenceGroupHopping ::= ENUMERATED
type DmrsUplinkconfigTransformprecodingenabledSequencegrouphopping struct {
	Value utils.ENUMERATED
}

const (
	DmrsUplinkconfigTransformprecodingenabledSequencegrouphoppingEnumeratedNothing = iota
	DmrsUplinkconfigTransformprecodingenabledSequencegrouphoppingEnumeratedDisabled
)

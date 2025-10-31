package ies

import "rrc/utils"

// DMRS-UplinkConfig-transformPrecodingEnabled-sequenceHopping ::= ENUMERATED
type DmrsUplinkconfigTransformprecodingenabledSequencehopping struct {
	Value utils.ENUMERATED
}

const (
	DmrsUplinkconfigTransformprecodingenabledSequencehoppingEnumeratedNothing = iota
	DmrsUplinkconfigTransformprecodingenabledSequencehoppingEnumeratedEnabled
)

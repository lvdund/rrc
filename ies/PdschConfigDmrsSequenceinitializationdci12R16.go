package ies

import "rrc/utils"

// PDSCH-Config-dmrs-SequenceInitializationDCI-1-2-r16 ::= ENUMERATED
type PdschConfigDmrsSequenceinitializationdci12R16 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigDmrsSequenceinitializationdci12R16EnumeratedNothing = iota
	PdschConfigDmrsSequenceinitializationdci12R16EnumeratedEnabled
)

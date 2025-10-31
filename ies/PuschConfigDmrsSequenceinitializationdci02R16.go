package ies

import "rrc/utils"

// PUSCH-Config-dmrs-SequenceInitializationDCI-0-2-r16 ::= ENUMERATED
type PuschConfigDmrsSequenceinitializationdci02R16 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigDmrsSequenceinitializationdci02R16EnumeratedNothing = iota
	PuschConfigDmrsSequenceinitializationdci02R16EnumeratedEnabled
)

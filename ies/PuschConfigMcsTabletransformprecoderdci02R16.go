package ies

import "rrc/utils"

// PUSCH-Config-mcs-TableTransformPrecoderDCI-0-2-r16 ::= ENUMERATED
type PuschConfigMcsTabletransformprecoderdci02R16 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigMcsTabletransformprecoderdci02R16EnumeratedNothing = iota
	PuschConfigMcsTabletransformprecoderdci02R16EnumeratedQam256
	PuschConfigMcsTabletransformprecoderdci02R16EnumeratedQam64lowse
)

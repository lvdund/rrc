package ies

import "rrc/utils"

// PDSCH-Config-vrb-ToPRB-InterleaverDCI-1-2-r16 ::= ENUMERATED
type PdschConfigVrbToprbInterleaverdci12R16 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigVrbToprbInterleaverdci12R16EnumeratedNothing = iota
	PdschConfigVrbToprbInterleaverdci12R16EnumeratedN2
	PdschConfigVrbToprbInterleaverdci12R16EnumeratedN4
)

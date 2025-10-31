package ies

import "rrc/utils"

// PDSCH-Config-vrb-ToPRB-Interleaver ::= ENUMERATED
type PdschConfigVrbToprbInterleaver struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigVrbToprbInterleaverEnumeratedNothing = iota
	PdschConfigVrbToprbInterleaverEnumeratedN2
	PdschConfigVrbToprbInterleaverEnumeratedN4
)

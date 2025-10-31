package ies

import "rrc/utils"

// SRS-Resource-ptrs-PortIndex ::= ENUMERATED
type SrsResourcePtrsPortindex struct {
	Value utils.ENUMERATED
}

const (
	SrsResourcePtrsPortindexEnumeratedNothing = iota
	SrsResourcePtrsPortindexEnumeratedN0
	SrsResourcePtrsPortindexEnumeratedN1
)

package ies

import "rrc/utils"

// LWA-Parameters-r13-lwa-BufferSize-r13 ::= ENUMERATED
type LwaParametersR13LwaBuffersizeR13 struct {
	Value utils.ENUMERATED
}

const (
	LwaParametersR13LwaBuffersizeR13EnumeratedNothing = iota
	LwaParametersR13LwaBuffersizeR13EnumeratedSupported
)

package ies

import "rrc/utils"

// LWA-Parameters-r13-lwa-r13 ::= ENUMERATED
type LwaParametersR13LwaR13 struct {
	Value utils.ENUMERATED
}

const (
	LwaParametersR13LwaR13EnumeratedNothing = iota
	LwaParametersR13LwaR13EnumeratedSupported
)

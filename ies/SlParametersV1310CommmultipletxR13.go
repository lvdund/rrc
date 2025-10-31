package ies

import "rrc/utils"

// SL-Parameters-v1310-commMultipleTx-r13 ::= ENUMERATED
type SlParametersV1310CommmultipletxR13 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1310CommmultipletxR13EnumeratedNothing = iota
	SlParametersV1310CommmultipletxR13EnumeratedSupported
)

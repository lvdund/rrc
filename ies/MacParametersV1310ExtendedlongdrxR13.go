package ies

import "rrc/utils"

// MAC-Parameters-v1310-extendedLongDRX-r13 ::= ENUMERATED
type MacParametersV1310ExtendedlongdrxR13 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1310ExtendedlongdrxR13EnumeratedNothing = iota
	MacParametersV1310ExtendedlongdrxR13EnumeratedSupported
)

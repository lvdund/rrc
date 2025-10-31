package ies

import "rrc/utils"

// MAC-Parameters-v1310-extendedMAC-LengthField-r13 ::= ENUMERATED
type MacParametersV1310ExtendedmacLengthfieldR13 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1310ExtendedmacLengthfieldR13EnumeratedNothing = iota
	MacParametersV1310ExtendedmacLengthfieldR13EnumeratedSupported
)

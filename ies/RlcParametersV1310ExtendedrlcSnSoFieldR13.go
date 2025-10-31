package ies

import "rrc/utils"

// RLC-Parameters-v1310-extendedRLC-SN-SO-Field-r13 ::= ENUMERATED
type RlcParametersV1310ExtendedrlcSnSoFieldR13 struct {
	Value utils.ENUMERATED
}

const (
	RlcParametersV1310ExtendedrlcSnSoFieldR13EnumeratedNothing = iota
	RlcParametersV1310ExtendedrlcSnSoFieldR13EnumeratedSupported
)

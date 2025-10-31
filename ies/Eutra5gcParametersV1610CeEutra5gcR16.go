package ies

import "rrc/utils"

// EUTRA-5GC-Parameters-v1610-ce-EUTRA-5GC-r16 ::= ENUMERATED
type Eutra5gcParametersV1610CeEutra5gcR16 struct {
	Value utils.ENUMERATED
}

const (
	Eutra5gcParametersV1610CeEutra5gcR16EnumeratedNothing = iota
	Eutra5gcParametersV1610CeEutra5gcR16EnumeratedSupported
)

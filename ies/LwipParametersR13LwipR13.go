package ies

import "rrc/utils"

// LWIP-Parameters-r13-lwip-r13 ::= ENUMERATED
type LwipParametersR13LwipR13 struct {
	Value utils.ENUMERATED
}

const (
	LwipParametersR13LwipR13EnumeratedNothing = iota
	LwipParametersR13LwipR13EnumeratedSupported
)

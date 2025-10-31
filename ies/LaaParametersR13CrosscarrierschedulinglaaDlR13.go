package ies

import "rrc/utils"

// LAA-Parameters-r13-crossCarrierSchedulingLAA-DL-r13 ::= ENUMERATED
type LaaParametersR13CrosscarrierschedulinglaaDlR13 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersR13CrosscarrierschedulinglaaDlR13EnumeratedNothing = iota
	LaaParametersR13CrosscarrierschedulinglaaDlR13EnumeratedSupported
)

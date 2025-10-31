package ies

import "rrc/utils"

// MUST-Parameters-r14-must-TM234-UpTo2Tx-r14 ::= ENUMERATED
type MustParametersR14MustTm234Upto2txR14 struct {
	Value utils.ENUMERATED
}

const (
	MustParametersR14MustTm234Upto2txR14EnumeratedNothing = iota
	MustParametersR14MustTm234Upto2txR14EnumeratedSupported
)

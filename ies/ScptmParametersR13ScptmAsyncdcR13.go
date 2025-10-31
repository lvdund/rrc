package ies

import "rrc/utils"

// SCPTM-Parameters-r13-scptm-AsyncDC-r13 ::= ENUMERATED
type ScptmParametersR13ScptmAsyncdcR13 struct {
	Value utils.ENUMERATED
}

const (
	ScptmParametersR13ScptmAsyncdcR13EnumeratedNothing = iota
	ScptmParametersR13ScptmAsyncdcR13EnumeratedSupported
)

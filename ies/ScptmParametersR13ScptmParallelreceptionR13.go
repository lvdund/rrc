package ies

import "rrc/utils"

// SCPTM-Parameters-r13-scptm-ParallelReception-r13 ::= ENUMERATED
type ScptmParametersR13ScptmParallelreceptionR13 struct {
	Value utils.ENUMERATED
}

const (
	ScptmParametersR13ScptmParallelreceptionR13EnumeratedNothing = iota
	ScptmParametersR13ScptmParallelreceptionR13EnumeratedSupported
)

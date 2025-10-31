package ies

import "rrc/utils"

// SCPTM-Parameters-r13-scptm-SCell-r13 ::= ENUMERATED
type ScptmParametersR13ScptmScellR13 struct {
	Value utils.ENUMERATED
}

const (
	ScptmParametersR13ScptmScellR13EnumeratedNothing = iota
	ScptmParametersR13ScptmScellR13EnumeratedSupported
)

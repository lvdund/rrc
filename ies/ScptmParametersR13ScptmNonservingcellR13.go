package ies

import "rrc/utils"

// SCPTM-Parameters-r13-scptm-NonServingCell-r13 ::= ENUMERATED
type ScptmParametersR13ScptmNonservingcellR13 struct {
	Value utils.ENUMERATED
}

const (
	ScptmParametersR13ScptmNonservingcellR13EnumeratedNothing = iota
	ScptmParametersR13ScptmNonservingcellR13EnumeratedSupported
)

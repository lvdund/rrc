package ies

import "rrc/utils"

// CA-ParametersNRDC-v1700-condPSCellAdditionNRDC-r17 ::= ENUMERATED
type CaParametersnrdcV1700CondpscelladditionnrdcR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrdcV1700CondpscelladditionnrdcR17EnumeratedNothing = iota
	CaParametersnrdcV1700CondpscelladditionnrdcR17EnumeratedSupported
)

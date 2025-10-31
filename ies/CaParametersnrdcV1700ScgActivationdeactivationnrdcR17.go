package ies

import "rrc/utils"

// CA-ParametersNRDC-v1700-scg-ActivationDeactivationNRDC-r17 ::= ENUMERATED
type CaParametersnrdcV1700ScgActivationdeactivationnrdcR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrdcV1700ScgActivationdeactivationnrdcR17EnumeratedNothing = iota
	CaParametersnrdcV1700ScgActivationdeactivationnrdcR17EnumeratedSupported
)

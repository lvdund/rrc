package ies

import "rrc/utils"

// CA-ParametersNRDC-v1700-scg-ActivationDeactivationResumeNRDC-r17 ::= ENUMERATED
type CaParametersnrdcV1700ScgActivationdeactivationresumenrdcR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrdcV1700ScgActivationdeactivationresumenrdcR17EnumeratedNothing = iota
	CaParametersnrdcV1700ScgActivationdeactivationresumenrdcR17EnumeratedSupported
)

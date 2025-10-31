package ies

import "rrc/utils"

// MRDC-Parameters-v1700-scg-ActivationDeactivationResumeENDC-r17 ::= ENUMERATED
type MrdcParametersV1700ScgActivationdeactivationresumeendcR17 struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersV1700ScgActivationdeactivationresumeendcR17EnumeratedNothing = iota
	MrdcParametersV1700ScgActivationdeactivationresumeendcR17EnumeratedSupported
)

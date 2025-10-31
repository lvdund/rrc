package ies

import "rrc/utils"

// MRDC-Parameters-v1700-scg-ActivationDeactivationENDC-r17 ::= ENUMERATED
type MrdcParametersV1700ScgActivationdeactivationendcR17 struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersV1700ScgActivationdeactivationendcR17EnumeratedNothing = iota
	MrdcParametersV1700ScgActivationdeactivationendcR17EnumeratedSupported
)

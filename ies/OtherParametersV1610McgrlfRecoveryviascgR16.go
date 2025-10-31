package ies

import "rrc/utils"

// Other-Parameters-v1610-mcgRLF-RecoveryViaSCG-r16 ::= ENUMERATED
type OtherParametersV1610McgrlfRecoveryviascgR16 struct {
	Value utils.ENUMERATED
}

const (
	OtherParametersV1610McgrlfRecoveryviascgR16EnumeratedNothing = iota
	OtherParametersV1610McgrlfRecoveryviascgR16EnumeratedSupported
)

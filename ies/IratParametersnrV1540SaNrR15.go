package ies

import "rrc/utils"

// IRAT-ParametersNR-v1540-sa-NR-r15 ::= ENUMERATED
type IratParametersnrV1540SaNrR15 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrV1540SaNrR15EnumeratedNothing = iota
	IratParametersnrV1540SaNrR15EnumeratedSupported
)

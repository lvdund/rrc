package ies

import "rrc/utils"

// IRAT-ParametersUTRA-v9c0-srvcc-FromUTRA-FDD-ToGERAN-r9 ::= ENUMERATED
type IratParametersutraV9c0SrvccFromutraFddTogeranR9 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersutraV9c0SrvccFromutraFddTogeranR9EnumeratedNothing = iota
	IratParametersutraV9c0SrvccFromutraFddTogeranR9EnumeratedSupported
)

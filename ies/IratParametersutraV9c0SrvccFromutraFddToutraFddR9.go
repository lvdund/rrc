package ies

import "rrc/utils"

// IRAT-ParametersUTRA-v9c0-srvcc-FromUTRA-FDD-ToUTRA-FDD-r9 ::= ENUMERATED
type IratParametersutraV9c0SrvccFromutraFddToutraFddR9 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersutraV9c0SrvccFromutraFddToutraFddR9EnumeratedNothing = iota
	IratParametersutraV9c0SrvccFromutraFddToutraFddR9EnumeratedSupported
)

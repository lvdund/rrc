package ies

import "rrc/utils"

// IRAT-ParametersUTRA-v9h0-mfbi-UTRA-r9 ::= ENUMERATED
type IratParametersutraV9h0MfbiUtraR9 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersutraV9h0MfbiUtraR9EnumeratedNothing = iota
	IratParametersutraV9h0MfbiUtraR9EnumeratedSupported
)

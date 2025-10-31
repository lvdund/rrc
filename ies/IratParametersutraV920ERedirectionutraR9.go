package ies

import "rrc/utils"

// IRAT-ParametersUTRA-v920-e-RedirectionUTRA-r9 ::= ENUMERATED
type IratParametersutraV920ERedirectionutraR9 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersutraV920ERedirectionutraR9EnumeratedNothing = iota
	IratParametersutraV920ERedirectionutraR9EnumeratedSupported
)

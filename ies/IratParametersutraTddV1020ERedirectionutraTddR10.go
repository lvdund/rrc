package ies

import "rrc/utils"

// IRAT-ParametersUTRA-TDD-v1020-e-RedirectionUTRA-TDD-r10 ::= ENUMERATED
type IratParametersutraTddV1020ERedirectionutraTddR10 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersutraTddV1020ERedirectionutraTddR10EnumeratedNothing = iota
	IratParametersutraTddV1020ERedirectionutraTddR10EnumeratedSupported
)

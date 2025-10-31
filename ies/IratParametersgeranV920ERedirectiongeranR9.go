package ies

import "rrc/utils"

// IRAT-ParametersGERAN-v920-e-RedirectionGERAN-r9 ::= ENUMERATED
type IratParametersgeranV920ERedirectiongeranR9 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersgeranV920ERedirectiongeranR9EnumeratedNothing = iota
	IratParametersgeranV920ERedirectiongeranR9EnumeratedSupported
)

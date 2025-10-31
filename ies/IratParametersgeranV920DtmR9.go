package ies

import "rrc/utils"

// IRAT-ParametersGERAN-v920-dtm-r9 ::= ENUMERATED
type IratParametersgeranV920DtmR9 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersgeranV920DtmR9EnumeratedNothing = iota
	IratParametersgeranV920DtmR9EnumeratedSupported
)

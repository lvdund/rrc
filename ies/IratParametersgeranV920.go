package ies

import "rrc/utils"

// IRAT-ParametersGERAN-v920 ::= SEQUENCE
type IratParametersgeranV920 struct {
	DtmR9               *utils.ENUMERATED
	ERedirectiongeranR9 *utils.ENUMERATED
}

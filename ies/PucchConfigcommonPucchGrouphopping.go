package ies

import "rrc/utils"

// PUCCH-ConfigCommon-pucch-GroupHopping ::= ENUMERATED
type PucchConfigcommonPucchGrouphopping struct {
	Value utils.ENUMERATED
}

const (
	PucchConfigcommonPucchGrouphoppingEnumeratedNothing = iota
	PucchConfigcommonPucchGrouphoppingEnumeratedNeither
	PucchConfigcommonPucchGrouphoppingEnumeratedEnable
	PucchConfigcommonPucchGrouphoppingEnumeratedDisable
)

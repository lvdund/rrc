package ies

import "rrc/utils"

// MBMS-CarrierType-r14-carrierType-r14 ::= ENUMERATED
type MbmsCarriertypeR14CarriertypeR14 struct {
	Value utils.ENUMERATED
}

const (
	MbmsCarriertypeR14CarriertypeR14EnumeratedNothing = iota
	MbmsCarriertypeR14CarriertypeR14EnumeratedMbms
	MbmsCarriertypeR14CarriertypeR14EnumeratedFembmsmixed
	MbmsCarriertypeR14CarriertypeR14EnumeratedFembmsdedicated
)

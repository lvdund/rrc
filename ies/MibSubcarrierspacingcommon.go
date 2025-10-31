package ies

import "rrc/utils"

// MIB-subCarrierSpacingCommon ::= ENUMERATED
type MibSubcarrierspacingcommon struct {
	Value utils.ENUMERATED
}

const (
	MibSubcarrierspacingcommonEnumeratedNothing = iota
	MibSubcarrierspacingcommonEnumeratedScs15or60
	MibSubcarrierspacingcommonEnumeratedScs30or120
)

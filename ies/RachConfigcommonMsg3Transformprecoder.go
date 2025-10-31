package ies

import "rrc/utils"

// RACH-ConfigCommon-msg3-transformPrecoder ::= ENUMERATED
type RachConfigcommonMsg3Transformprecoder struct {
	Value utils.ENUMERATED
}

const (
	RachConfigcommonMsg3TransformprecoderEnumeratedNothing = iota
	RachConfigcommonMsg3TransformprecoderEnumeratedEnabled
)

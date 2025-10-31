package ies

import "rrc/utils"

// MsgA-PUSCH-Config-r16-msgA-TransformPrecoder-r16 ::= ENUMERATED
type MsgaPuschConfigR16MsgaTransformprecoderR16 struct {
	Value utils.ENUMERATED
}

const (
	MsgaPuschConfigR16MsgaTransformprecoderR16EnumeratedNothing = iota
	MsgaPuschConfigR16MsgaTransformprecoderR16EnumeratedEnabled
	MsgaPuschConfigR16MsgaTransformprecoderR16EnumeratedDisabled
)

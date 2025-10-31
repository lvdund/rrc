package ies

import "rrc/utils"

// MAC-Parameters-v1610-rai-SupportEnh-r16 ::= ENUMERATED
type MacParametersV1610RaiSupportenhR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1610RaiSupportenhR16EnumeratedNothing = iota
	MacParametersV1610RaiSupportenhR16EnumeratedSupported
)

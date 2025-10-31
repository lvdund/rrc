package ies

import "rrc/utils"

// MAC-Parameters-NB-v1610-rai-SupportEnh-r16 ::= ENUMERATED
type MacParametersNbV1610RaiSupportenhR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersNbV1610RaiSupportenhR16EnumeratedNothing = iota
	MacParametersNbV1610RaiSupportenhR16EnumeratedSupported
)

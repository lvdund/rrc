package ies

import "rrc/utils"

// MAC-Parameters-v1440-rai-Support-r14 ::= ENUMERATED
type MacParametersV1440RaiSupportR14 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1440RaiSupportR14EnumeratedNothing = iota
	MacParametersV1440RaiSupportR14EnumeratedSupported
)

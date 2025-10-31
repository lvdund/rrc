package ies

import "rrc/utils"

// MAC-Parameters-NB-r14-rai-Support-r14 ::= ENUMERATED
type MacParametersNbR14RaiSupportR14 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersNbR14RaiSupportR14EnumeratedNothing = iota
	MacParametersNbR14RaiSupportR14EnumeratedSupported
)

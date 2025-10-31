package ies

import "rrc/utils"

// MAC-ParametersCommon-lcid-ExtensionIAB-r16 ::= ENUMERATED
type MacParameterscommonLcidExtensioniabR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonLcidExtensioniabR16EnumeratedNothing = iota
	MacParameterscommonLcidExtensioniabR16EnumeratedSupported
)

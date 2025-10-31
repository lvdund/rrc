package ies

import "rrc/utils"

// MAC-ParametersCommon-lch-ToConfiguredGrantMapping-r16 ::= ENUMERATED
type MacParameterscommonLchToconfiguredgrantmappingR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonLchToconfiguredgrantmappingR16EnumeratedNothing = iota
	MacParameterscommonLchToconfiguredgrantmappingR16EnumeratedSupported
)

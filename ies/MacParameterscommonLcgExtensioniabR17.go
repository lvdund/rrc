package ies

import "rrc/utils"

// MAC-ParametersCommon-lcg-ExtensionIAB-r17 ::= ENUMERATED
type MacParameterscommonLcgExtensioniabR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonLcgExtensioniabR17EnumeratedNothing = iota
	MacParameterscommonLcgExtensioniabR17EnumeratedSupported
)

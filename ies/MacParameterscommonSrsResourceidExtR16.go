package ies

import "rrc/utils"

// MAC-ParametersCommon-srs-ResourceId-Ext-r16 ::= ENUMERATED
type MacParameterscommonSrsResourceidExtR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonSrsResourceidExtR16EnumeratedNothing = iota
	MacParameterscommonSrsResourceidExtR16EnumeratedSupported
)

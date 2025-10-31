package ies

import "rrc/utils"

// MAC-ParametersCommon-singlePHR-P-r16 ::= ENUMERATED
type MacParameterscommonSinglephrPR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonSinglephrPR16EnumeratedNothing = iota
	MacParameterscommonSinglephrPR16EnumeratedSupported
)

package ies

import "rrc/utils"

// MAC-ParametersCommon-spCell-BFR-CBRA-r16 ::= ENUMERATED
type MacParameterscommonSpcellBfrCbraR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonSpcellBfrCbraR16EnumeratedNothing = iota
	MacParameterscommonSpcellBfrCbraR16EnumeratedSupported
)

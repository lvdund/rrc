package ies

import "rrc/utils"

// MAC-ParametersCommon-dummy ::= ENUMERATED
type MacParameterscommonDummy struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonDummyEnumeratedNothing = iota
	MacParameterscommonDummyEnumeratedSupported
)

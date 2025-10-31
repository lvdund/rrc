package ies

import "rrc/utils"

// MAC-Parameters-v1430-skipUplinkDynamic-r14 ::= ENUMERATED
type MacParametersV1430SkipuplinkdynamicR14 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1430SkipuplinkdynamicR14EnumeratedNothing = iota
	MacParametersV1430SkipuplinkdynamicR14EnumeratedSupported
)

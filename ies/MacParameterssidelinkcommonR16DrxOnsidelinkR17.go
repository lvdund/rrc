package ies

import "rrc/utils"

// MAC-ParametersSidelinkCommon-r16-drx-OnSidelink-r17 ::= ENUMERATED
type MacParameterssidelinkcommonR16DrxOnsidelinkR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterssidelinkcommonR16DrxOnsidelinkR17EnumeratedNothing = iota
	MacParameterssidelinkcommonR16DrxOnsidelinkR17EnumeratedSupported
)

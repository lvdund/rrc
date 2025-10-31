package ies

import "rrc/utils"

// MAC-ParametersSidelink-r17-drx-OnSidelink-r17 ::= ENUMERATED
type MacParameterssidelinkR17DrxOnsidelinkR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterssidelinkR17DrxOnsidelinkR17EnumeratedNothing = iota
	MacParameterssidelinkR17DrxOnsidelinkR17EnumeratedSupported
)

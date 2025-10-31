package ies

import "rrc/utils"

// MAC-ParametersSidelinkCommon-r16-lcp-RestrictionSidelink-r16 ::= ENUMERATED
type MacParameterssidelinkcommonR16LcpRestrictionsidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterssidelinkcommonR16LcpRestrictionsidelinkR16EnumeratedNothing = iota
	MacParameterssidelinkcommonR16LcpRestrictionsidelinkR16EnumeratedSupported
)

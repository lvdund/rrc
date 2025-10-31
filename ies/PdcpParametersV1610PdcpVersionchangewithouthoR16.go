package ies

import "rrc/utils"

// PDCP-Parameters-v1610-pdcp-VersionChangeWithoutHO-r16 ::= ENUMERATED
type PdcpParametersV1610PdcpVersionchangewithouthoR16 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersV1610PdcpVersionchangewithouthoR16EnumeratedNothing = iota
	PdcpParametersV1610PdcpVersionchangewithouthoR16EnumeratedSupported
)

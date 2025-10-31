package ies

import "rrc/utils"

// MAC-ParametersCommon-lcp-Restriction ::= ENUMERATED
type MacParameterscommonLcpRestriction struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonLcpRestrictionEnumeratedNothing = iota
	MacParameterscommonLcpRestrictionEnumeratedSupported
)

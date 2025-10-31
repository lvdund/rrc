package ies

import "rrc/utils"

// MAC-Parameters-v1550-eLCID-Support-r15 ::= ENUMERATED
type MacParametersV1550ElcidSupportR15 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersV1550ElcidSupportR15EnumeratedNothing = iota
	MacParametersV1550ElcidSupportR15EnumeratedSupported
)

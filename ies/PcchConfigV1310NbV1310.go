package ies

import "rrc/utils"

// PCCH-Config-v1310-nB-v1310 ::= ENUMERATED
type PcchConfigV1310NbV1310 struct {
	Value utils.ENUMERATED
}

const (
	PcchConfigV1310NbV1310EnumeratedNothing = iota
	PcchConfigV1310NbV1310EnumeratedOne64tht
	PcchConfigV1310NbV1310EnumeratedOne128tht
	PcchConfigV1310NbV1310EnumeratedOne256tht
)

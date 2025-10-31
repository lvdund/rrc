package ies

import "rrc/utils"

// SPS-Config-mcs-Table ::= ENUMERATED
type SpsConfigMcsTable struct {
	Value utils.ENUMERATED
}

const (
	SpsConfigMcsTableEnumeratedNothing = iota
	SpsConfigMcsTableEnumeratedQam64lowse
)

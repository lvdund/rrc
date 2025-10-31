package ies

import "rrc/utils"

// CE-PUSCH-MultiTB-Config-r16-interleaving-r16 ::= ENUMERATED
type CePuschMultitbConfigR16InterleavingR16 struct {
	Value utils.ENUMERATED
}

const (
	CePuschMultitbConfigR16InterleavingR16EnumeratedNothing = iota
	CePuschMultitbConfigR16InterleavingR16EnumeratedOn
)

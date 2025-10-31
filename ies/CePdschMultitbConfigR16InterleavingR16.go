package ies

import "rrc/utils"

// CE-PDSCH-MultiTB-Config-r16-interleaving-r16 ::= ENUMERATED
type CePdschMultitbConfigR16InterleavingR16 struct {
	Value utils.ENUMERATED
}

const (
	CePdschMultitbConfigR16InterleavingR16EnumeratedNothing = iota
	CePdschMultitbConfigR16InterleavingR16EnumeratedOn
)

package ies

import "rrc/utils"

// SL-HoppingConfigDisc-r12-c-r12 ::= ENUMERATED
type SlHoppingconfigdiscR12CR12 struct {
	Value utils.ENUMERATED
}

const (
	SlHoppingconfigdiscR12CR12EnumeratedNothing = iota
	SlHoppingconfigdiscR12CR12EnumeratedN1
	SlHoppingconfigdiscR12CR12EnumeratedN5
)

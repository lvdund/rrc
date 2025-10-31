package ies

import "rrc/utils"

// SL-PTRS-Config-r16-sl-PTRS-RE-Offset-r16 ::= ENUMERATED
type SlPtrsConfigR16SlPtrsReOffsetR16 struct {
	Value utils.ENUMERATED
}

const (
	SlPtrsConfigR16SlPtrsReOffsetR16EnumeratedNothing = iota
	SlPtrsConfigR16SlPtrsReOffsetR16EnumeratedOffset01
	SlPtrsConfigR16SlPtrsReOffsetR16EnumeratedOffset10
	SlPtrsConfigR16SlPtrsReOffsetR16EnumeratedOffset11
)

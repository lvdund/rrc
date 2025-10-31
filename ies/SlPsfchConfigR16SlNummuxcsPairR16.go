package ies

import "rrc/utils"

// SL-PSFCH-Config-r16-sl-NumMuxCS-Pair-r16 ::= ENUMERATED
type SlPsfchConfigR16SlNummuxcsPairR16 struct {
	Value utils.ENUMERATED
}

const (
	SlPsfchConfigR16SlNummuxcsPairR16EnumeratedNothing = iota
	SlPsfchConfigR16SlNummuxcsPairR16EnumeratedN1
	SlPsfchConfigR16SlNummuxcsPairR16EnumeratedN2
	SlPsfchConfigR16SlNummuxcsPairR16EnumeratedN3
	SlPsfchConfigR16SlNummuxcsPairR16EnumeratedN6
)

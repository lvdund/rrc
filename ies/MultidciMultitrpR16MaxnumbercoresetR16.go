package ies

import "rrc/utils"

// MultiDCI-MultiTRP-r16-maxNumberCORESET-r16 ::= ENUMERATED
type MultidciMultitrpR16MaxnumbercoresetR16 struct {
	Value utils.ENUMERATED
}

const (
	MultidciMultitrpR16MaxnumbercoresetR16EnumeratedNothing = iota
	MultidciMultitrpR16MaxnumbercoresetR16EnumeratedN2
	MultidciMultitrpR16MaxnumbercoresetR16EnumeratedN3
	MultidciMultitrpR16MaxnumbercoresetR16EnumeratedN4
	MultidciMultitrpR16MaxnumbercoresetR16EnumeratedN5
)

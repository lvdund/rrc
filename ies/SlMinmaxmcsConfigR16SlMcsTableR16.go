package ies

import "rrc/utils"

// SL-MinMaxMCS-Config-r16-sl-MCS-Table-r16 ::= ENUMERATED
type SlMinmaxmcsConfigR16SlMcsTableR16 struct {
	Value utils.ENUMERATED
}

const (
	SlMinmaxmcsConfigR16SlMcsTableR16EnumeratedNothing = iota
	SlMinmaxmcsConfigR16SlMcsTableR16EnumeratedQam64
	SlMinmaxmcsConfigR16SlMcsTableR16EnumeratedQam256
	SlMinmaxmcsConfigR16SlMcsTableR16EnumeratedQam64lowse
)

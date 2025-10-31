package ies

import "rrc/utils"

// SL-MinMaxMCS-Config-r16 ::= SEQUENCE
type SlMinmaxmcsConfigR16 struct {
	SlMcsTableR16    SlMinmaxmcsConfigR16SlMcsTableR16
	SlMinmcsPsschR16 utils.INTEGER `lb:0,ub:27`
	SlMaxmcsPsschR16 utils.INTEGER `lb:0,ub:31`
}

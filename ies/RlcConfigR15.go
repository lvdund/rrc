package ies

import "rrc/utils"

// RLC-Config-r15 ::= SEQUENCE
// Extensible
type RlcConfigR15 struct {
	ModeR15                  interface{}
	ReestablishrlcR15        *utils.ENUMERATED
	RlcOutoforderdeliveryR15 *utils.ENUMERATED
}

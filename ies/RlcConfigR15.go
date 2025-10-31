package ies

// RLC-Config-r15 ::= SEQUENCE
// Extensible
type RlcConfigR15 struct {
	ModeR15                  RlcConfigR15ModeR15
	ReestablishrlcR15        *bool
	RlcOutoforderdeliveryR15 *bool
}

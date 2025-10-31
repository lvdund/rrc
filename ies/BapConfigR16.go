package ies

import "rrc/utils"

// BAP-Config-r16 ::= SEQUENCE
// Extensible
type BapConfigR16 struct {
	BapAddressR16              *utils.BITSTRING `lb:10,ub:10`
	DefaultulBapRoutingidR16   *BapRoutingidR16
	DefaultulBhRlcChannelR16   *BhRlcChannelidR16
	FlowcontrolfeedbacktypeR16 *BapConfigR16FlowcontrolfeedbacktypeR16
}

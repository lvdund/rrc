package ies

import "rrc/utils"

// BAP-Config-r16-flowControlFeedbackType-r16 ::= ENUMERATED
type BapConfigR16FlowcontrolfeedbacktypeR16 struct {
	Value utils.ENUMERATED
}

const (
	BapConfigR16FlowcontrolfeedbacktypeR16EnumeratedNothing = iota
	BapConfigR16FlowcontrolfeedbacktypeR16EnumeratedPerbh_Rlc_Channel
	BapConfigR16FlowcontrolfeedbacktypeR16EnumeratedPerroutingid
	BapConfigR16FlowcontrolfeedbacktypeR16EnumeratedBoth
)

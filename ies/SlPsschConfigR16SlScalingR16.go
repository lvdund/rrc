package ies

import "rrc/utils"

// SL-PSSCH-Config-r16-sl-Scaling-r16 ::= ENUMERATED
type SlPsschConfigR16SlScalingR16 struct {
	Value utils.ENUMERATED
}

const (
	SlPsschConfigR16SlScalingR16EnumeratedNothing = iota
	SlPsschConfigR16SlScalingR16EnumeratedF0p5
	SlPsschConfigR16SlScalingR16EnumeratedF0p65
	SlPsschConfigR16SlScalingR16EnumeratedF0p8
	SlPsschConfigR16SlScalingR16EnumeratedF1
)

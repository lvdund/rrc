package ies

import "rrc/utils"

// AUL-Config-r15-setup-transmissionModeUL-AUL-r15 ::= ENUMERATED
type AulConfigR15SetupTransmissionmodeulAulR15 struct {
	Value utils.ENUMERATED
}

const (
	AulConfigR15SetupTransmissionmodeulAulR15EnumeratedNothing = iota
	AulConfigR15SetupTransmissionmodeulAulR15EnumeratedTm1
	AulConfigR15SetupTransmissionmodeulAulR15EnumeratedTm2
)

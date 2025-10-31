package ies

import "rrc/utils"

// AUL-Config-r15-setup-contentionWindowSizeTimer-r15 ::= ENUMERATED
type AulConfigR15SetupContentionwindowsizetimerR15 struct {
	Value utils.ENUMERATED
}

const (
	AulConfigR15SetupContentionwindowsizetimerR15EnumeratedNothing = iota
	AulConfigR15SetupContentionwindowsizetimerR15EnumeratedN0
	AulConfigR15SetupContentionwindowsizetimerR15EnumeratedN5
	AulConfigR15SetupContentionwindowsizetimerR15EnumeratedN10
)

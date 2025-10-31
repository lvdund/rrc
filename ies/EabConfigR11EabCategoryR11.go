package ies

import "rrc/utils"

// EAB-Config-r11-eab-Category-r11 ::= ENUMERATED
type EabConfigR11EabCategoryR11 struct {
	Value utils.ENUMERATED
}

const (
	EabConfigR11EabCategoryR11EnumeratedNothing = iota
	EabConfigR11EabCategoryR11EnumeratedA
	EabConfigR11EabCategoryR11EnumeratedB
	EabConfigR11EabCategoryR11EnumeratedC
)

package ies

import "rrc/utils"

// Alpha-r12 ::= ENUMERATED
type AlphaR12 struct {
	Value utils.ENUMERATED
}

const (
	AlphaR12Al0  = 0
	AlphaR12Al04 = 1
	AlphaR12Al05 = 2
	AlphaR12Al06 = 3
	AlphaR12Al07 = 4
	AlphaR12Al08 = 5
	AlphaR12Al09 = 6
	AlphaR12Al1  = 7
)

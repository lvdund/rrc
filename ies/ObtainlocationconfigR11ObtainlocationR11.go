package ies

import "rrc/utils"

// ObtainLocationConfig-r11-obtainLocation-r11 ::= ENUMERATED
type ObtainlocationconfigR11ObtainlocationR11 struct {
	Value utils.ENUMERATED
}

const (
	ObtainlocationconfigR11ObtainlocationR11EnumeratedNothing = iota
	ObtainlocationconfigR11ObtainlocationR11EnumeratedSetup
)

package ies

import "rrc/utils"

// UAC-BarringPerPLMN-List-r15 ::= SEQUENCE OF UAC-BarringPerPLMN-r15
// SIZE (1.. maxPLMN-r11)
type UacBarringperplmnListR15 struct {
	Value []UacBarringperplmnR15
}

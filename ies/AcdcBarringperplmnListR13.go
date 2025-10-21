package ies

import "rrc/utils"

// ACDC-BarringPerPLMN-List-r13 ::= SEQUENCE OF ACDC-BarringPerPLMN-r13
// SIZE (1.. maxPLMN-r11)
type AcdcBarringperplmnListR13 struct {
	Value utils.Sequence[AcdcBarringperplmnR13]
}

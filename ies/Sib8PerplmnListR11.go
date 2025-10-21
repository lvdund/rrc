package ies

import "rrc/utils"

// SIB8-PerPLMN-List-r11 ::= SEQUENCE OF SIB8-PerPLMN-r11
// SIZE (1..maxPLMN-r11)
type Sib8PerplmnListR11 struct {
	Value utils.Sequence[Sib8PerplmnR11]
}

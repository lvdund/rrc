package ies

import "rrc/utils"

// UDT-RestrictingPerPLMN-List-r13 ::= SEQUENCE OF UDT-RestrictingPerPLMN-r13
// SIZE (1..maxPLMN-r11)
type UdtRestrictingperplmnListR13 struct {
	Value []UdtRestrictingperplmnR13
}

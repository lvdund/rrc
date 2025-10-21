package ies

import "rrc/utils"

// MBMS-SAI-List-r11 ::= SEQUENCE OF MBMS-SAI-r11
// SIZE (1..maxSAI-MBMS-r11)
type MbmsSaiListR11 struct {
	Value utils.Sequence[MbmsSaiR11]
}

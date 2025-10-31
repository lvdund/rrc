package ies

// MBMS-SAI-List-r11 ::= SEQUENCE OF MBMS-SAI-r11
// SIZE (1..maxSAI-MBMS-r11)
type MbmsSaiListR11 struct {
	Value []MbmsSaiR11 `lb:1,ub:maxSAIMbmsR11`
}

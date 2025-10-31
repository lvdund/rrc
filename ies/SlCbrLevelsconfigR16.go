package ies

// SL-CBR-LevelsConfig-r16 ::= SEQUENCE OF SL-CBR-r16
// SIZE (1..maxCBR-Level-r16)
type SlCbrLevelsconfigR16 struct {
	Value []SlCbrR16 `lb:1,ub:maxCBRLevelR16`
}

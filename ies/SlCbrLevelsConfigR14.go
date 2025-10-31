package ies

// SL-CBR-Levels-Config-r14 ::= SEQUENCE OF SL-CBR-r14
// SIZE (1..maxCBR-Level-r14)
type SlCbrLevelsConfigR14 struct {
	Value []SlCbrR14 `lb:1,ub:maxCBRLevelR14`
}

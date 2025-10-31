package ies

// CSI-IM-ConfigToAddModList-r11 ::= SEQUENCE OF CSI-IM-Config-r11
// SIZE (1..maxCSI-IM-r11)
type CsiImConfigtoaddmodlistR11 struct {
	Value []CsiImConfigR11 `lb:1,ub:maxCSIImR11`
}

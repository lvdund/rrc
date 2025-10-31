package ies

// CSI-IM-ConfigToReleaseList-r11 ::= SEQUENCE OF CSI-IM-ConfigId-r11
// SIZE (1..maxCSI-IM-r11)
type CsiImConfigtoreleaselistR11 struct {
	Value []CsiImConfigidR11 `lb:1,ub:maxCSIImR11`
}

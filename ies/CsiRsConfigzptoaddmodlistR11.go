package ies

// CSI-RS-ConfigZPToAddModList-r11 ::= SEQUENCE OF CSI-RS-ConfigZP-r11
// SIZE (1..maxCSI-RS-ZP-r11)
type CsiRsConfigzptoaddmodlistR11 struct {
	Value []CsiRsConfigzpR11 `lb:1,ub:maxCSIRsZpR11`
}

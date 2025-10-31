package ies

// CSI-RS-ConfigNZPToAddModList-r11 ::= SEQUENCE OF CSI-RS-ConfigNZP-r11
// SIZE (1..maxCSI-RS-NZP-r11)
type CsiRsConfignzptoaddmodlistR11 struct {
	Value []CsiRsConfignzpR11 `lb:1,ub:maxCSIRsNzpR11`
}

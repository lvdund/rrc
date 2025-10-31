package ies

// CSI-RS-ConfigNZPToAddModList-r15 ::= SEQUENCE OF CSI-RS-ConfigNZP-r11
// SIZE (1..maxCSI-RS-NZP-r13)
type CsiRsConfignzptoaddmodlistR15 struct {
	Value []CsiRsConfignzpR11 `lb:1,ub:maxCSIRsNzpR13`
}

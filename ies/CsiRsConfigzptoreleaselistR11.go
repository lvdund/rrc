package ies

// CSI-RS-ConfigZPToReleaseList-r11 ::= SEQUENCE OF CSI-RS-ConfigZPId-r11
// SIZE (1..maxCSI-RS-ZP-r11)
type CsiRsConfigzptoreleaselistR11 struct {
	Value []CsiRsConfigzpidR11 `lb:1,ub:maxCSIRsZpR11`
}

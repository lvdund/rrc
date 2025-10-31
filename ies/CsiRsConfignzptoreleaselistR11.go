package ies

// CSI-RS-ConfigNZPToReleaseList-r11 ::= SEQUENCE OF CSI-RS-ConfigNZPId-r11
// SIZE (1..maxCSI-RS-NZP-r11)
type CsiRsConfignzptoreleaselistR11 struct {
	Value []CsiRsConfignzpidR11 `lb:1,ub:maxCSIRsNzpR11`
}

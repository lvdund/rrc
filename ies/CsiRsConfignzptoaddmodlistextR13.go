package ies

// CSI-RS-ConfigNZPToAddModListExt-r13 ::= SEQUENCE OF CSI-RS-ConfigNZP-r11
// SIZE (1..maxCSI-RS-NZP-v1310)
type CsiRsConfignzptoaddmodlistextR13 struct {
	Value []CsiRsConfignzpR11 `lb:1,ub:maxCSIRsNzpV1310`
}

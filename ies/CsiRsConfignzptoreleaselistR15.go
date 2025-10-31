package ies

// CSI-RS-ConfigNZPToReleaseList-r15 ::= SEQUENCE OF CSI-RS-ConfigNZPId-r13
// SIZE (1..maxCSI-RS-NZP-r13)
type CsiRsConfignzptoreleaselistR15 struct {
	Value []CsiRsConfignzpidR13 `lb:1,ub:maxCSIRsNzpR13`
}

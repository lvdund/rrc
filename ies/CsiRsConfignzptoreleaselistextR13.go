package ies

// CSI-RS-ConfigNZPToReleaseListExt-r13 ::= SEQUENCE OF CSI-RS-ConfigNZPId-v1310
// SIZE (1..maxCSI-RS-NZP-v1310)
type CsiRsConfignzptoreleaselistextR13 struct {
	Value []CsiRsConfignzpidV1310 `lb:1,ub:maxCSIRsNzpV1310`
}

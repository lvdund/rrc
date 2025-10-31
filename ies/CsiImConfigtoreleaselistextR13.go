package ies

// CSI-IM-ConfigToReleaseListExt-r13 ::= SEQUENCE OF CSI-IM-ConfigId-v1310
// SIZE (1..maxCSI-IM-v1310)
type CsiImConfigtoreleaselistextR13 struct {
	Value []CsiImConfigidV1310 `lb:1,ub:maxCSIImV1310`
}

package ies

// CSI-IM-ConfigToAddModListExt-r13 ::= SEQUENCE OF CSI-IM-ConfigExt-r12
// SIZE (1..maxCSI-IM-v1310)
type CsiImConfigtoaddmodlistextR13 struct {
	Value []CsiImConfigextR12 `lb:1,ub:maxCSIImV1310`
}

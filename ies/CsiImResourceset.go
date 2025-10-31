package ies

// CSI-IM-ResourceSet ::= SEQUENCE
// Extensible
type CsiImResourceset struct {
	CsiImResourcesetid CsiImResourcesetid
	CsiImResources     []CsiImResourceid `lb:1,ub:maxNrofCSIImResourcesperset`
}

package ies

// CSI-SSB-ResourceSet ::= SEQUENCE
// Extensible
type CsiSsbResourceset struct {
	CsiSsbResourcesetid         CsiSsbResourcesetid
	CsiSsbResourcelist          []SsbIndex                      `lb:1,ub:maxNrofCSISsbResourceperset`
	ServingadditionalpcilistR17 *[]ServingadditionalpciindexR17 `lb:1,ub:maxNrofCSISsbResourceperset,ext`
}

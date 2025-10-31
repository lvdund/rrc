package ies

// MultiPDSCH-TDRA-List-r17 ::= SEQUENCE OF MultiPDSCH-TDRA-r17
// SIZE (1.. maxNrofDL-AllocationsExt-r17)
type MultipdschTdraListR17 struct {
	Value []MultipdschTdraR17 `lb:1,ub:maxNrofDLAllocationsextR17`
}

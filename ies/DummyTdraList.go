package ies

// Dummy-TDRA-List ::= SEQUENCE OF MultiPDSCH-TDRA-r17
// SIZE (1.. maxNrofDL-Allocations)
type DummyTdraList struct {
	Value []MultipdschTdraR17 `lb:1,ub:maxNrofDLAllocations`
}

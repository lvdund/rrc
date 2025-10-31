package ies

// MeasResultServCellListSCG-Ext-r13 ::= SEQUENCE OF MeasResultServCellSCG-r12
// SIZE (1..maxServCell-r13)
type MeasresultservcelllistscgExtR13 struct {
	Value []MeasresultservcellscgR12 `lb:1,ub:maxServCellR13`
}

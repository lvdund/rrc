package ies

// MeasResultCSI-RS-List-r12 ::= SEQUENCE OF MeasResultCSI-RS-r12
// SIZE (1..maxCellReport)
type MeasresultcsiRsListR12 struct {
	Value []MeasresultcsiRsR12 `lb:1,ub:maxCellReport`
}

package ies

// MeasResultIdle-r15 ::= SEQUENCE
// Extensible
type MeasresultidleR15 struct {
	MeasresultservingcellR15 MeasresultidleR15MeasresultservingcellR15
	MeasresultneighcellsR15  *MeasresultidleR15MeasresultneighcellsR15
}

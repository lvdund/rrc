package ies

// MeasResultCellNR-r15 ::= SEQUENCE
// Extensible
type MeasresultcellnrR15 struct {
	PciR15                   PhyscellidnrR15
	MeasresultcellR15        MeasresultnrR15
	MeasresultrsIndexlistR15 *MeasresultssbIndexlistR15
}

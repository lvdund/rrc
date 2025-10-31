package ies

// MeasResultIdle-r15-measResultNeighCells-r15 ::= CHOICE
// Extensible
const (
	MeasresultidleR15MeasresultneighcellsR15ChoiceNothing = iota
	MeasresultidleR15MeasresultneighcellsR15ChoiceMeasresultidlelisteutraR15
)

type MeasresultidleR15MeasresultneighcellsR15 struct {
	Choice                     uint64
	MeasresultidlelisteutraR15 *MeasresultidlelisteutraR15
}

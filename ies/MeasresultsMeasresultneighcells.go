package ies

// MeasResults-measResultNeighCells ::= CHOICE
// Extensible
const (
	MeasresultsMeasresultneighcellsChoiceNothing = iota
	MeasresultsMeasresultneighcellsChoiceMeasresultlisteutra
	MeasresultsMeasresultneighcellsChoiceMeasresultlistutra
	MeasresultsMeasresultneighcellsChoiceMeasresultlistgeran
	MeasresultsMeasresultneighcellsChoiceMeasresultscdma2000
	MeasresultsMeasresultneighcellsChoiceMeasresultneighcelllistnrR15
)

type MeasresultsMeasresultneighcells struct {
	Choice                       uint64
	Measresultlisteutra          *Measresultlisteutra
	Measresultlistutra           *Measresultlistutra
	Measresultlistgeran          *Measresultlistgeran
	Measresultscdma2000          *Measresultscdma2000
	MeasresultneighcelllistnrR15 *MeasresultcelllistnrR15
}

package ies

// RLF-Report-r9-measResultNeighCells-r9 ::= SEQUENCE
type RlfReportR9MeasresultneighcellsR9 struct {
	MeasresultlisteutraR9 *Measresultlist2eutraR9
	MeasresultlistutraR9  *Measresultlist2utraR9
	MeasresultlistgeranR9 *Measresultlistgeran
	Measresultscdma2000R9 *Measresultlist2cdma2000R9
}

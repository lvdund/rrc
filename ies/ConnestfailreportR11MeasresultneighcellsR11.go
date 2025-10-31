package ies

// ConnEstFailReport-r11-measResultNeighCells-r11 ::= SEQUENCE
type ConnestfailreportR11MeasresultneighcellsR11 struct {
	MeasresultlisteutraR11 *Measresultlist2eutraR9
	MeasresultlistutraR11  *Measresultlist2utraR9
	MeasresultlistgeranR11 *Measresultlistgeran
	Measresultscdma2000R11 *Measresultlist2cdma2000R9
}

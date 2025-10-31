package ies

// LogMeasInfo-r10-measResultNeighCells-r10 ::= SEQUENCE
type LogmeasinfoR10MeasresultneighcellsR10 struct {
	MeasresultlisteutraR10    *Measresultlist2eutraR9
	MeasresultlistutraR10     *Measresultlist2utraR9
	MeasresultlistgeranR10    *Measresultlist2geranR10
	Measresultlistcdma2000R10 *Measresultlist2cdma2000R9
}

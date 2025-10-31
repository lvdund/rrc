package ies

// NeighCellCDMA2000 ::= SEQUENCE
type Neighcellcdma2000 struct {
	Bandclass             Bandclasscdma2000
	Neighcellsperfreqlist Neighcellsperbandclasslistcdma2000
}

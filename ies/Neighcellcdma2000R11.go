package ies

// NeighCellCDMA2000-r11 ::= SEQUENCE
type Neighcellcdma2000R11 struct {
	Bandclass            Bandclasscdma2000
	NeighfreqinfolistR11 []Neighcellsperbandclasscdma2000R11 `lb:1,ub:16`
}

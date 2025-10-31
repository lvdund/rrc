package ies

// NeighCellsPerBandclassCDMA2000-r11 ::= SEQUENCE
type Neighcellsperbandclasscdma2000R11 struct {
	Arfcn             ArfcnValuecdma2000
	PhyscellidlistR11 []Physcellidcdma2000 `lb:1,ub:40`
}

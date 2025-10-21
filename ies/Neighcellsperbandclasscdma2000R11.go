package ies

import "rrc/utils"

// NeighCellsPerBandclassCDMA2000-r11 ::= SEQUENCE
type Neighcellsperbandclasscdma2000R11 struct {
	Arfcn             ArfcnValuecdma2000
	PhyscellidlistR11 interface{}
}

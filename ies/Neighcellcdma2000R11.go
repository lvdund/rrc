package ies

import "rrc/utils"

// NeighCellCDMA2000-r11 ::= SEQUENCE
type Neighcellcdma2000R11 struct {
	Bandclass            Bandclasscdma2000
	NeighfreqinfolistR11 interface{}
}

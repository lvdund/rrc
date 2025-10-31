package ies

import "rrc/utils"

// BandClassInfoCDMA2000 ::= SEQUENCE
// Extensible
type Bandclassinfocdma2000 struct {
	Bandclass               Bandclasscdma2000
	Cellreselectionpriority *Cellreselectionpriority
	ThreshxHigh             utils.INTEGER `lb:0,ub:63`
	ThreshxLow              utils.INTEGER `lb:0,ub:63`
}
